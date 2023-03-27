package generate

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func GRawTransactions(numberOfTransactions int, wFileName string, tFileName string) {
	// 打开包含私钥和地址的文件
	file, err := os.Open(wFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建名为 RawTransactions.csv 的文件，并写入表头
	transactionsFile, err := os.Create(tFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer transactionsFile.Close()

	transactionsWriter := csv.NewWriter(transactionsFile)
	defer transactionsWriter.Flush()

	// 从文件中读取每个钱包的私钥和地址，并为每个钱包创建一笔交易
	reader := csv.NewReader(file)
	for i := 0; i < numberOfTransactions; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		privateKey, err := crypto.HexToECDSA(record[0])
		if err != nil {
			log.Fatal(err)
		}

		address := common.HexToAddress(record[1])
		nonce, gasPrice, chainID, err := getNonce(address)
		if err != nil {
			log.Fatal(err)
		}

		// 构造交易数据
		gasLimit := uint64(21000) // 固定的 gasLimit

		to := common.HexToAddress("0x1234567890123456789012345678901234567890") // 接收方地址

		var data []byte
		tx := types.NewTransaction(nonce, to, big.NewInt(1), gasLimit, gasPrice, data)

		// 使用钱包的私钥进行签名
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}

		txb, err := rlp.EncodeToBytes(signedTx)
		if err != nil {
			log.Fatalf("encode error: %v", err)
		}

		record = []string{common.Bytes2Hex(txb)}
		if err := transactionsWriter.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("%d 笔交易已生成并写入文件！", numberOfTransactions)
}

func getNonce(address common.Address) (uint64, *big.Int, *big.Int, error) {
	client, err := ethclient.Dial("http://127.0.0.1:9545")
	if err != nil {
		return 0, nil, nil, err
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, nil, nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return 0, nil, nil, err
	}

	return nonce, gasPrice, chainID, nil
}
