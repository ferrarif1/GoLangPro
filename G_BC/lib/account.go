package lib

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsContract(c *ethclient.Client, a common.Address) (bool, error) {

	// nil is latest block
	bytecode, err := c.CodeAt(context.Background(), a, nil)
	if err != nil {
		return false, err
	}

	isContract := len(bytecode) > 0

	return isContract, nil
}

func BalanceAt(c *ethclient.Client, a common.Address, bn int64) (balance *big.Int, err error) {

	if bn == -1 {
		balance, err = c.BalanceAt(context.Background(), a, nil)
	} else {
		blockNumber := big.NewInt(bn)
		balance, err = c.BalanceAt(context.Background(), a, blockNumber)
	}

	return
}

func PendingBalanceAt(c *ethclient.Client, a string) (pendingBalance *big.Int, err error) {

	account := common.HexToAddress(a)
	pendingBalance, err = c.PendingBalanceAt(context.Background(), account)

	return
}

func PendingNonceAt(c *ethclient.Client, a common.Address) (nonce uint64, err error) {

	nonce, err = c.PendingNonceAt(context.Background(), a)

	return
}
