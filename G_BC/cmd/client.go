package cmd

import (
	"context"
	"fmt"
	"mantle/test/config"
	"mantle/test/lib"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:     "c",
	Aliases: []string{"client"},
	Short:   "服务于mantle rpc client的测试工具",
	Long:    "服务于mantle rpc client的测试工具",
	RunE: func(cmd *cobra.Command, args []string) error {

		ci, err := c.ChainID(context.Background())
		fmt.Println("ChainID: ", ci)

		return err
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

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

var url string
var tokenAddress string
var c *ethclient.Client
var b int64
var bh string

func init() {
	clientCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "")

	clientCmd.PersistentFlags().StringVar(&tokenAddress, "tokenAddress", "", "token Contract address")
	clientCmd.PersistentFlags().Int64VarP(&b, "blockNum", "b", -1, "区块高度")
	clientCmd.PersistentFlags().StringVar(&bh, "blockHash", "0", "区块hash")

	mtCmd.AddCommand(clientCmd)
}
