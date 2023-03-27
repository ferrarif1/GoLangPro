package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"mantle/test/config"
	"mantle/test/lib"
)

var (
	prv        string
	from       string
	to         string
	value      int64
	dataString string
)

func init() {
	clientCmd.AddCommand(ctCmd)
	ctCmd.AddCommand(ethCmd)

	ctCmd.PersistentFlags().StringP("toAddress", "t", "", "receiver address")
	ctCmd.PersistentFlags().StringP("privateKey", "p", "", "privateKey of sender address")

	ctCmd.PersistentFlags().Int64VarP(&value, "value", "v", 1, "transfer value")
	ctCmd.PersistentFlags().StringVarP(&dataString, "dataString", "d", "", "transfer message")

	//将配置文件和入参绑定，使用入参替换默认配置
	config.AppConfig.BindPFlag("privateKey", ctCmd.PersistentFlags().Lookup("privateKey"))
	config.AppConfig.BindPFlag("to", ctCmd.PersistentFlags().Lookup("to"))
}

var ctCmd = &cobra.Command{
	Use:     "t",
	Aliases: []string{"transfer"},
	Short:   "执行transfer相关操作",
	Long:    "执行transfer相关操作",
	Example: "mt client transfer SendETHTransaction",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("nothing todo……")
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		prv = config.AppConfig.GetString("privateKey")
		to = config.AppConfig.GetString("to")

		_, _, fromAddress, err := lib.AnalysePrivateKey(prv)
		if err != nil {
			return err
		}

		from = fromAddress.Hex()
		return nil
	},
}

var ethCmd = &cobra.Command{
	Use:     "e",
	Aliases: []string{"SendETHTransaction"},
	Short:   "查询最新区块高度",
	Long:    "查询最新区块高度",
	RunE: func(cmd *cobra.Command, args []string) error {

		tx, cid, privateKey, err := lib.SignETHTx1(c, prv, from, to, value, []byte(dataString))
		if err != nil {
			return err
		}

		signedTx, err := lib.SignETHTx2(tx, cid, privateKey)
		if err != nil {
			return err
		}

		txHash, err := lib.SendTransaction(c, signedTx)
		if err != nil {
			return err
		}

		fmt.Println("txHash: ", txHash)
		return nil
	},
	PreRunE: ctCmd.PreRunE,
}
