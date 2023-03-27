package lib

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func NewEthClient(u string) (c *ethclient.Client, e error) {

	// 连接节点
	c, e = ethclient.Dial(u)
	return
}

func GenerateEOA(prik string) (privateKey *ecdsa.PrivateKey,
	publicKey *ecdsa.PublicKey, address common.Address, err error) {

	if prik == "" {
		// 生成椭圆曲线
		curve := elliptic.P256()

		// 生成私钥
		privateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
	} else {
		privateKey, err = crypto.HexToECDSA(prik)
	}

	pubk := privateKey.Public()
	publicKey, ok := pubk.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address = crypto.PubkeyToAddress(*publicKey)

	return
}

func QBasic(c *ethclient.Client, address common.Address) (nonce uint64,
	gasPrice *big.Int, chainID *big.Int, err error) {

	nonce, err = c.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, nil, nil, err
	}

	gasPrice, err = c.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, nil, nil, err
	}

	chainID, err = c.NetworkID(context.Background())

	return
}

func newETHTransaction(c *ethclient.Client, fromAddress common.Address, toAddress string,
	value int64, gasLimit uint64, data []byte) (tx *types.Transaction, chainID *big.Int,
	err error) {

	nonce, gasPrice, chainID, err := QBasic(c, fromAddress)
	if err != nil {
		return nil, nil, err
	}

	if value == 0 {
		value = 1000000000000000000
	}
	v := big.NewInt(value)

	if gasLimit == 0 {
		gasLimit = 21000
	}

	to := common.HexToAddress(toAddress)
	tx = types.NewTransaction(nonce, to, v, gasLimit, gasPrice, data)

	// gas, _ := c.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	From:     fromAddress,
	// 	To:       &to,
	// 	Value:    v,
	// 	Gas:      gasLimit,
	// 	GasPrice: gasPrice,
	// })
	// fmt.Println("gas: ", gas)

	return
}

func newERC20Transaction(c *ethclient.Client, tokenAddress string, fromAddress common.Address, toAddress string,
	value int64, gasLimit uint64, data []byte) (tx *types.Transaction, chainID *big.Int,
	err error) {

	nonce, gasPrice, chainID, err := QBasic(c, fromAddress)
	if err != nil {
		return nil, nil, err
	}

	if value == 0 {
		value = 1000000000000000000
	}
	v := big.NewInt(value)

	if gasLimit == 0 {
		gasLimit = 21000
	}

	to := common.HexToAddress(toAddress)

	tokenAd := common.HexToAddress("tokenAddress")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(to.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	paddedAmount := common.LeftPadBytes(v.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err = c.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAd,
		Data: data,
	})
	fmt.Println("gasLimit: ", gasLimit)

	if err != nil {
		return nil, nil, err
	}

	tx = types.NewTransaction(nonce, to, big.NewInt(0), gasLimit, gasPrice, data)

	return
}

func SignedTx(c *ethclient.Client, prik string, toAddress string, tokenAddress string,
	value int64, gasLimit uint64, data []byte) (signedTx *types.Transaction,
	err error) {

	privateKey, _, fromAddress, err := GenerateEOA(prik)
	if err != nil {
		return nil, err
	}

	var tx *types.Transaction
	var cid *big.Int
	if tokenAddress == "" {
		tx, cid, err = newETHTransaction(c, fromAddress, toAddress, value, gasLimit, data)
		if err != nil {
			return nil, err
		}
	} else {
		tx, cid, err = newERC20Transaction(c, tokenAddress, fromAddress, toAddress, value, gasLimit, data)
		if err != nil {
			return nil, err
		}
	}

	signedTx, err = types.SignTx(tx, types.NewEIP155Signer(cid), privateKey)

	return
}
