package cmd

import (
	"errors"
	"fmt"
	"mantle/test/config"
	"mantle/test/generate"
	"mantle/test/lib"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:       "g",
	Aliases:   []string{"generate"},
	ValidArgs: []string{"w", "t", "e"},
	Args:      cobra.OnlyValidArgs,
	Short:     "生成钱包、序列化交易以及反序列化交易并发起请求",
	Long:      "生成钱包、序列化交易以及反序列化交易并发起请求",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 || len(args) > 1 {
			fmt.Println(errors.New("只能输入一个参数! "))
			// fmt.Println("need args: w \\ t \\ r")
			// fmt.Println("w: generate wallets")
			// fmt.Println("t: generate RawTransactions")
			// fmt.Println("e: execute RawTransactions")
			return
		}

		switch args[0] {
		case "w":
			generate.GWallets(numberOfWR, wfileName)
		case "t":
			generate.GRawTransactions(numberOfWR, wfileName, rfileName)
		case "e":
			fmt.Println("todo……")
		}

	},

	PreRun: func(cmd *cobra.Command, args []string) {

		// 创建rpc链接
		var err error
		if url == "" {
			url = config.AppConfig.GetString("url")
		}
		c, err = lib.NewEthClient(url)
		if err != nil {
			panic(err)
		}
	},
}

var numberOfWR int
var wfileName string
var rfileName string

func init() {
	generateCmd.Flags().IntVarP(&numberOfWR, "numberOfWR", "n", 10, "生成的wallets数量, 也就是RawTransactions数量")
	generateCmd.Flags().StringVarP(&wfileName, "wfileName", "w", "wallets.csv", "保存wallets的文件")
	generateCmd.Flags().StringVarP(&rfileName, "rfileName", "r", "RawTransactions.csv", "保存RawTransactions的文件")

	mtCmd.AddCommand(generateCmd)
}
