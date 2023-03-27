package cmd

import (
	"errors"
	"fmt"
	"mantle/test/lib"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	clientCmd.AddCommand(txCmd)
	txCmd.AddCommand(rCmd)
}

var txCmd = &cobra.Command{
	Use:     "tx",
	Aliases: []string{"transaction"},
	Args:    cobra.ExactArgs(1),
	Short:   "执行transaction相关操作",
	Long:    "执行transaction相关操作",
	Example: "mt client transaction <txHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		tx, _, err := lib.TransactionByHash(c, args[0])
		if err != nil {
			return err
		}
		fmt.Printf("根据txHash: %v 查询到txGas: %v\n", args[0], tx.Gas())
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {
			return errors.New("txkHash非法……")
		}
		return nil
	},
}

var rCmd = &cobra.Command{
	Use:     "r",
	Aliases: []string{"receipt"},
	Args:    cobra.ExactArgs(1),
	Short:   "执行transaction相关操作",
	Long:    "执行transaction相关操作",
	Example: "mt client transaction receipt <txHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		tr, err := lib.TransactionReceipt(c, args[0])
		if err != nil {
			return err
		}
		fmt.Printf("根据txHash: %v 查询到Receipt Status: %v\n", args[0], tr.Status)
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {
			return errors.New("txkHash非法……")
		}
		return nil
	},
}
