package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"mantle/test/lib"
)

var (
	address   common.Address
	isAddress bool
)

func init() {
	clientCmd.AddCommand(caCmd)
	caCmd.AddCommand(icCmd)
	caCmd.AddCommand(bCmd)
	caCmd.AddCommand(tbCmd)
}

var caCmd = &cobra.Command{
	Use:     "a",
	Aliases: []string{"account"},
	Args:    cobra.ExactArgs(1),
	Short:   "执行account相关操作",
	Long:    "执行account相关操作",
	Example: "mt client account <address>",
	RunE: func(cmd *cobra.Command, args []string) error {

		ci, err := c.ChainID(context.Background())
		fmt.Println("ChainID: ", ci)

		return err
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		address, isAddress = lib.CheckAddress(args[0])
		if !isAddress {
			return errors.New("输入的address非法,请确认……")
		}

		return nil
	},
}

var icCmd = &cobra.Command{
	Use:     "ic",
	Aliases: []string{"isContract"},
	Args:    caCmd.Args,
	Short:   "判断地址是否为合约地址",
	Long:    "判断地址是否为合约地址",
	Example: "mt client account isContract <address>",
	RunE: func(cmd *cobra.Command, args []string) error {

		isContract, err := lib.IsContract(c, address)

		fmt.Printf("地址: %v 是否为合约: %v\n", args[0], isContract)
		return err
	},
	PreRunE: caCmd.PreRunE,
}

var bCmd = &cobra.Command{
	Use:     "b",
	Aliases: []string{"balance"},
	Short:   "查询balance",
	Long:    "查询balance",
	RunE: func(cmd *cobra.Command, args []string) error {

		balance, err := lib.BalanceAt(c, address, b)
		if err != nil {
			return err
		}

		if b == -1 {
			fmt.Printf("%v 在最新区块的balance: %v\n", address.Hex(), balance)
		} else {
			fmt.Printf("%v 在区块: %v 的balance: %v\n", args[0], b, balance)
		}
		return nil
	},
	PreRunE: caCmd.PreRunE,
}

var tbCmd = &cobra.Command{
	Use:     "tb",
	Aliases: []string{"token balance"},
	Short:   "查询token 的 balance",
	Long:    "查询token 的 balance",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("to do……")
		return nil
	},
}
