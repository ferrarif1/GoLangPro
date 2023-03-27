package lib

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BlockNumber(c *ethclient.Client) (bn uint64, err error) {

	bn, err = c.BlockNumber(context.Background())
	if err != nil {
		return bn, err
	}

	return
}

func HeaderByNumber(c *ethclient.Client, bn int64) (header *types.Header, err error) {

	blockNumber := big.NewInt(bn)
	if bn == -1 {
		blockNumber = nil
	}

	header, err = c.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return header, err
	}

	return
}

func HeaderByHash(c *ethclient.Client, bh string) (header *types.Header, err error) {

	blockHash := common.HexToHash(bh)
	header, err = c.HeaderByHash(context.Background(), blockHash)
	if err != nil {
		return header, err
	}

	return
}

func BlockByNumber(c *ethclient.Client, bn int64) (block *types.Block, err error) {

	blockNumber := big.NewInt(bn)
	block, err = c.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return block, err
	}

	return
}

func BlockByHash(c *ethclient.Client, bh string) (block *types.Block, err error) {

	blockHash := common.HexToHash(bh)
	block, err = c.BlockByHash(context.Background(), blockHash)
	if err != nil {
		return block, err
	}

	return
}

func TransactionCount(c *ethclient.Client, bh string) (count uint, err error) {

	block, err := BlockByHash(c, bh)
	if err != nil {
		return count, err
	}

	count, err = c.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		return count, err
	}

	return
}
