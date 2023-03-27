package lib

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SignETHTx1(c *ethclient.Client, prv string, from string, to string, v int64, data []byte) (
	tx *types.Transaction, chainID *big.Int, privateKey *ecdsa.PrivateKey, err error) {

	value := big.NewInt(v)
	gasLimit := uint64(21000)

	privateKey, err = crypto.HexToECDSA(prv)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress(from)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err = c.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(to)
	tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	return
}

func SignETHTx2(tx *types.Transaction, chainID *big.Int, prv *ecdsa.PrivateKey) (
	signedTx *types.Transaction, err error) {

	signedTx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), prv)
	return
}

func SendTransaction(c *ethclient.Client, signedTx *types.Transaction) (txHash string, err error) {
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}

	txHash = signedTx.Hash().Hex()
	return
}
