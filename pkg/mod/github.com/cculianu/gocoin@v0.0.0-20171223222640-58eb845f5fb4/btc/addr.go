package btc

import (
	"bytes"
	"fmt"
	"math/big"
	"errors"
)

/*
const (
	ADDRVER_BTC = byte(0x00)
	ADDRVER_TESTNET = byte(0x6F)
)
*/

type BtcAddr struct {
	Version byte
	Hash160 [20]byte
	Checksum []byte
	Pubkey []byte
	Enc58str string

	// This is normally not used, unless for GetAllUnspent() purposes
	Extra struct {
		Label string
		Wallet string
		Virgin bool
	}
}

func NewAddrFromString(hs string) (a *BtcAddr, e error) {
	dec := Decodeb58(hs)
	if dec == nil {
		e = errors.New("Cannot decode b58 string *"+hs+"*")
		return
	}
	if (len(dec)<25) {
		dec = append(bytes.Repeat([]byte{0}, 25-len(dec)), dec...)
	}
	if (len(dec)==25) {
		sh := Sha2Sum(dec[0:21])
		if !bytes.Equal(sh[:4], dec[21:25]) {
			e = errors.New("Address Checksum error")
		} else {
			a = new(BtcAddr)
			a.Version = dec[0]
			copy(a.Hash160[:], dec[1:21])
			a.Checksum = make([]byte, 4)
			copy(a.Checksum, dec[21:25])
			a.Enc58str = hs
		}
	} else {
		e = errors.New(fmt.Sprintf("Unsupported hash length %d", len(dec)))
	}
	return
}

func NewAddrFromHash160(in []byte, ver byte) (a *BtcAddr) {
	a = new(BtcAddr)
	a.Version = ver
	copy(a.Hash160[:], in[:])
	return
}


func NewAddrFromP2SH(in []byte, ver byte) (a *BtcAddr) {
	a = new(BtcAddr)
	a.Version = ver
	copy(a.Hash160[:], in[:])
	return
}


func NewAddrFromPubkey(in []byte, ver byte) (a *BtcAddr) {
	a = new(BtcAddr)
	a.Pubkey = make([]byte, len(in))
	copy(a.Pubkey[:], in[:])
	a.Version = ver
	RimpHash(in, a.Hash160[:])
	return
}


func AddrVerPubkey(testnet bool) byte {
	if testnet {
		return 111
	} else {
		return 0
	}
}


func AddrVerScript(testnet bool) byte {
	if testnet {
		return 196
	} else {
		return 5
	}
}


func NewAddrFromPkScript(scr []byte, testnet bool) (*BtcAddr) {
	if len(scr)==25 && scr[0]==0x76 && scr[1]==0xa9 && scr[2]==0x14 && scr[23]==0x88 && scr[24]==0xac {
		return NewAddrFromHash160(scr[3:23], AddrVerPubkey(testnet))
	} else if len(scr)==67 && scr[0]==0x41 && scr[66]==0xac {
		return NewAddrFromPubkey(scr[1:66], AddrVerPubkey(testnet))
	} else if len(scr)==35 && scr[0]==0x21 && scr[34]==0xac {
		return NewAddrFromPubkey(scr[1:34], AddrVerPubkey(testnet))
	} else if len(scr)==23 && scr[0]==0xa9 && scr[1]==0x14 && scr[22]==0x87 {
		return NewAddrFromHash160(scr[2:22], AddrVerScript(testnet))
	}
	return nil
}


func NewAddrFromDataWithSum(in []byte, ver byte) (a *BtcAddr, e error) {
	var ad [25]byte
	ad[0] = ver
	copy(ad[1:25], in[:])
	sh := Sha2Sum(ad[0:21])
	if !bytes.Equal(in[20:24], sh[:4]) {
		e = errors.New("Address Checksum error")
		return
	}

	copy(ad[21:25], sh[:4])

	a = new(BtcAddr)
	a.Version = ver
	copy(a.Hash160[:], in[:])

	a.Checksum = make([]byte, 4)
	copy(a.Checksum, sh[:4])
	return
}

// Base58 encoded address
func (a *BtcAddr) String() string {
	if a.Enc58str=="" {
		var ad [25]byte
		ad[0] = a.Version
		copy(ad[1:21], a.Hash160[:])
		if a.Checksum==nil {
			sh := Sha2Sum(ad[0:21])
			a.Checksum = make([]byte, 4)
			copy(a.Checksum, sh[:4])
		}
		copy(ad[21:25], a.Checksum[:])
		a.Enc58str = Encodeb58(ad[:])
	}
	return a.Enc58str
}

// String with a label
func (a *BtcAddr) StringLab() (s string) {
	s = a.String()

	if a.Extra.Wallet!="" {
		s += " " + a.Extra.Wallet + ":"
	}
	if a.Extra.Label!="" {
		s += " " + a.Extra.Label
	}
	if a.Extra.Virgin {
		s += " ***"
	}
	return
}

// Check if a pk_script send coins to this address
func (a *BtcAddr) Owns(scr []byte) (yes bool) {
	// The most common spend script
	if len(scr)==25 && scr[0]==0x76 && scr[1]==0xa9 && scr[2]==0x14 && scr[23]==0x88 && scr[24]==0xac {
		yes = bytes.Equal(scr[3:23], a.Hash160[:])
		return
	}

	// Spend script with an entire public key
	if len(scr)==67 && scr[0]==0x41 && scr[1]==0x04 && scr[66]==0xac {
		if a.Pubkey == nil {
			h := Rimp160AfterSha256(scr[1:66])
			if h == a.Hash160 {
				a.Pubkey = make([]byte, 65)
				copy(a.Pubkey, scr[1:66])
				yes = true
			}
			return
		}
		yes = bytes.Equal(scr[1:34], a.Pubkey[:33])
		return
	}

	// Spend script with a compressed public key
	if len(scr)==35 && scr[0]==0x21 && (scr[1]==0x02 || scr[1]==0x03) && scr[34]==0xac {
		if a.Pubkey == nil {
			h := Rimp160AfterSha256(scr[1:34])
			if h == a.Hash160 {
				a.Pubkey = make([]byte, 33)
				copy(a.Pubkey, scr[1:34])
				yes = true
			}
			return
		}
		yes = bytes.Equal(scr[1:34], a.Pubkey[:33])
		return
	}

	return
}

/*
	Just for information:

	// P2SH transaction
	if len(scr)==23 && scr[0]==0xa9 && scr[1]==0x14 && scr[22]==0x87 {
		return
	}

	// Escrow
	if len(scr)==201 && scr[0]==0x51 && scr[1]==0x41 && scr[199]==0x53 && scr[200]==0xae {
		return
	}
*/


func (a *BtcAddr) OutScript() (res []byte) {
	res = make([]byte, 25)
	res[0] = 0x76
	res[1] = 0xa9
	res[2] = 20
	copy(res[3:23], a.Hash160[:])
	res[23] = 0x88
	res[24] = 0xac
	return
}

var b58set []byte = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func b58chr2int(chr byte) int {
	for i:=range b58set {
		if b58set[i]==chr {
			return i
		}
	}
	return -1
}


var bn0 *big.Int = big.NewInt(0)
var bn58 *big.Int = big.NewInt(58)

func Encodeb58(a []byte) (s string) {
	idx := len(a) * 138 / 100 + 1
	buf := make([]byte, idx)
	bn := new(big.Int).SetBytes(a)
	var mo *big.Int
	for bn.Cmp(bn0) != 0 {
		bn, mo = bn.DivMod(bn, bn58, new(big.Int))
		idx--
		buf[idx] = b58set[mo.Int64()]
	}
	for i := range a {
		if a[i]!=0 {
			break
		}
		idx--
		buf[idx] = b58set[0]
	}

	s = string(buf[idx:])

	return
}

func Decodeb58(s string) []byte {
	bn := big.NewInt(0)
	for i := range s {
		v := b58chr2int(byte(s[i]))
		if v < 0 {
			return nil
		}
		bn = bn.Mul(bn, bn58)
		bn = bn.Add(bn, big.NewInt(int64(v)))
	}
	return bn.Bytes()
}
