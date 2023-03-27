package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func maina() {
	// client, err := ethclient.Dial("https://rpc.testnet.mantle.xyz")

	var wg1 = new(sync.WaitGroup)
	for i := 0; i < 100; i++ {

		wg1.Add(1)
		client, err := ethclient.Dial("https://rpc0.testnet.mantle.xyz")
		if err != nil {
			log.Fatal(err)
		}

		go func(a int, wg *sync.WaitGroup) {

			defer wg.Done()
			for j := 0; j < 5800; j++ {

				blockNumber := big.NewInt(int64(j + a*5800))
				block, err := client.BlockByNumber(context.Background(), blockNumber)
				if err != nil {
					fmt.Println("get block err, err: ", err, "blockNumber: ", blockNumber)
					break
				}

				blockd := block
				if a > 0 {
					blockd, err = client.BlockByNumber(context.Background(), big.NewInt(int64(j+a*5800)-1))
					if err != nil {
						fmt.Println("get blockd err, err: ", err, "blockNumber: ", big.NewInt(int64(j+a*5800)-1))
						break
					}
				}

				fmt.Println("Block  Nm:", block.Number())

				block1, err := client.BlockByHash(context.Background(), block.Hash())
				if err != nil {
					fmt.Println("get BlockByHash err, err: ", err, "BlockHash: ", block.Hash())
					break
				}

				if block1.Number().Cmp(block.Number()) != 0 {

					fmt.Printf(
						"Block 不一致, Block  Nm: %v, BlockHash: %v, Block1 Nm: %v, Block1Hash: %v\n",
						block.Number(), block.Hash(), block1.Number(), block1.Hash(),
					)
				}

				if a > 0 {
					if block1.ParentHash().Hex() != blockd.Hash().Hex() {

						fmt.Printf(
							"block1 ParentHash 不一致, Block1 Nm: %v, Blockd Nm: %v, ParentHash: %v, BlockdHash: %v\n",
							block1.Number(), blockd.Number(), block.ParentHash(), blockd.Hash(),
						)
					}

					if block.ParentHash().Hex() != blockd.Hash().Hex() {

						fmt.Printf(
							"block ParentHash 不一致, Block  Nm: %v, Blockd Nm: %v, ParentHash: %v, BlockdHash: %v\n",
							block.Number(), blockd.Number(), block.ParentHash(), blockd.Hash(),
						)
					}
				}
			}
		}(i, wg1)

	}
	wg1.Wait()

}

func main() {
	urls := []string{
		"http://10.45.25.148:8545",
		"http://10.45.26.68:8545",
		"http://10.45.27.102:8545",
		"http://10.45.24.37:8545",
		"http://10.45.27.191:8545",
		"http://10.45.26.205:8545",
		"http://10.45.24.231:8545",
		"http://10.45.27.193:8545",
		"http://10.45.27.124:8545",
	}

	clients := []*ethclient.Client{}
	for _, u := range urls {
		c, err := ethclient.Dial(u)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, c)
	}

	clients_len := len(clients)
	var wg1 = new(sync.WaitGroup)
	for i := 0; i < 100; i++ {

		wg1.Add(1)
		go func(a int, wg *sync.WaitGroup) {

			defer wg.Done()
			for j := 0; j < 5800; j++ {

				blocks := []*types.Block{}
				blockNumber := big.NewInt(int64(j + a*5800))

				for k := 0; k < clients_len; k++ {
					b, err := clients[k].BlockByNumber(context.Background(), blockNumber)
					if err != nil {
						fmt.Println("get block err, err: ", err, "blockNumber: ", blockNumber)
						break
					}
					blocks = append(blocks, b)
				}

				for l := 1; l < clients_len; l++ {
					r := compareBlock(blocks[0], blocks[l])
					if len(r) > 2 {
						fmt.Println(urls[0], urls[l], r)
					} else {
						fmt.Println("blockNumber: ", blockNumber, "finish!")
					}
				}

				sleep(500)
			}
		}(i, wg1)

	}

	wg1.Wait()
}

func compareBlock(b1 *types.Block, b2 *types.Block) []string {

	res := []string{b1.Number().String(), b2.Number().String()}

	if !bytes.Equal(b1.Extra(), b2.Extra()) {
		res = append(res, fmt.Sprintf("different extraData, b1: %v, b2: %v",
			hex.EncodeToString(b1.Extra()), hex.EncodeToString(b2.Extra())))
	}

	if b1.GasLimit() != b2.GasLimit() {
		res = append(res, "different GasLimit")
	}

	if b1.GasUsed() != b2.GasUsed() {
		res = append(res, "different GasUsed")
	}

	if b1.Hash().Hex() != b2.Hash().Hex() {
		res = append(res, "different Hash")
	}

	if b1.Difficulty().Cmp(b2.Difficulty()) != 0 {
		res = append(res, "different Difficulty")
	}

	if b1.Nonce() != b2.Nonce() {
		res = append(res, "different Nonce")
	}

	if b1.ParentHash().Hex() != b2.ParentHash().Hex() {
		res = append(res, "different ParentHash")
	}

	if b1.MixDigest().Hex() != b2.MixDigest().Hex() {
		res = append(res, "different MixDigest")
	}

	if b1.UncleHash().Hex() != b2.UncleHash().Hex() {
		res = append(res, "different UncleHash")
	}

	if b1.ReceiptHash().Hex() != b2.ReceiptHash().Hex() {
		res = append(res, "different ReceiptHash")
	}

	if b1.Root().Hex() != b2.Root().Hex() {
		res = append(res, "different Root")
	}

	if b1.Size() != b1.Size() {
		res = append(res, "different Size")
	}

	if b1.Time() != b2.Time() {
		res = append(res, "different Time")
	}

	if b1.Bloom() != b2.Bloom() {
		res = append(res, "different Bloom")
	}

	if b1.Coinbase().Hex() != b2.Coinbase().Hex() {
		res = append(res, "different miner")
	}

	return res
}

func sleep(ms int) {
	s := rand.Intn(ms)
	time.Sleep(time.Millisecond * time.Duration(s))
}
