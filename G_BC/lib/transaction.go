package lib

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransactionByHash(c *ethclient.Client, th string) (tx *types.Transaction, isPending bool, err error) {

	txHash := common.HexToHash(th)
	tx, isPending, err = c.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return
	}

	return
}

func TransactionReceipt(c *ethclient.Client, th string) (tr *types.Receipt, err error) {

	txHash := common.HexToHash(th)
	tr, err = c.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return
	}

	return
}
