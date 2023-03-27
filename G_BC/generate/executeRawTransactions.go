package generate

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func executeRawTransactions() {
	// 打开包含原始交易数据的文件
	file, err := os.Open("RawTransactions.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 从文件中读取每条原始交易数据，并提交交易
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		var tx types.Transaction
		err = rlp.DecodeBytes(common.Hex2Bytes(record[0]), &tx)

		from, err := types.Sender(types.HomesteadSigner{}, &tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(from)

		fmt.Println(tx)

		// 提交交易
		if err := SubmitTransaction(&tx); err != nil {
			log.Printf("交易 %s 提交失败：%v", tx.Hash().Hex(), err)
		} else {
			log.Printf("交易 %s 提交成功！", tx.Hash().Hex())
		}
	}
}

func SubmitTransaction(tx *types.Transaction) error {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}

	// // 获取交易的签名
	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	return err
	// }

	// signer := types.NewEIP155Signer(chainID)
	// signedTx, err := tx.WithSignature(signer, crypto.PubkeyToAddress(tx.From()))
	// if err != nil {
	// 	return err
	// }

	// // 将交易发送到以太坊网络
	// err = client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	return err
	// }

	return nil
}
