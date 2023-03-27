package lib

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CheckAddress(address string) (a common.Address, b bool) {

	if !common.IsHexAddress(address) {
		return a, b
	}

	return common.HexToAddress(address), true
}

func AnalysePrivateKey(prv string) (p *ecdsa.PrivateKey, publicKeyECDSA *ecdsa.PublicKey,
	fromAddress common.Address, err error) {

	privateKey, err := crypto.HexToECDSA(prv)
	if err != nil {
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}
