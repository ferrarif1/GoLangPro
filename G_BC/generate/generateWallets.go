package generate

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/csv"
	"encoding/hex"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

func GWallets(numberOfWallets int, fileName string) {
	// 创建名为 secret.csv 的文件，并写入表头
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 生成指定数量的钱包地址和私钥，并将它们写入文件
	for i := 0; i < numberOfWallets; i++ {
		// 生成一个新的私钥
		privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			log.Fatal(err)
		}

		// 将私钥转换为字节序列
		privateKeyBytes := privateKey.D.Bytes()

		// 将字节序列转换为十六进制字符串
		privateKeyHex := hex.EncodeToString(privateKeyBytes)

		// 根据私钥生成公钥和地址
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("生成公钥失败！")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		// 将私钥和地址写入文件
		record := []string{privateKeyHex, address}
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("%d 个钱包地址和私钥已生成并写入文件！", numberOfWallets)
}
