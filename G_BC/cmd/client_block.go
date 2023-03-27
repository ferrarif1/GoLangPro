package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"mantle/test/lib"
)

func init() {
	clientCmd.AddCommand(cbCmd)
	clientCmd.AddCommand(bnCmd)
	cbCmd.AddCommand(bhCmd)
	cbCmd.AddCommand(tcCmd)
}

var cbCmd = &cobra.Command{
	Use:     "b",
	Aliases: []string{"block"},
	Args:    cobra.ExactArgs(1),
	Short:   "执行block相关操作",
	Long:    "执行block相关操作",
	Example: "mt client block <blockNum | blockHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {

			blockNum, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			block, err := lib.BlockByNumber(c, blockNum)
			if err != nil {
				return err
			}
			fmt.Printf("根据num: %v 查询到了区块, block.timestamp: %v\n", blockNum, block.Time())

		} else {
			block, err := lib.BlockByHash(c, args[0])
			if err != nil {
				return err
			}

			fmt.Printf("根据blockhash: %v 查询到了区块, block.timestamp: %v\n", args[0], block.Time())
		}

		return nil
	},
}

var bnCmd = &cobra.Command{
	Use:     "bn",
	Aliases: []string{"blockNum"},
	Short:   "查询最新区块高度",
	Long:    "查询最新区块高度",
	RunE: func(cmd *cobra.Command, args []string) error {

		bn, err := lib.BlockNumber(c)
		if err != nil {
			return err
		}

		fmt.Println("最新区块高度为: ", bn)
		return nil
	},
}

var bhCmd = &cobra.Command{
	Use:     "h",
	Aliases: []string{"header"},
	Args:    cobra.ExactArgs(1),
	Short:   "查询指定区块的header",
	Long:    "查询指定区块的header",
	Example: "mt client block header <blockNum | blockHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {

			blockNum, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			header, err := lib.HeaderByNumber(c, blockNum)
			if err != nil {
				return err
			}
			fmt.Printf("区块 %v 的header.ParentHash为: %v\n", header.Number, header.ParentHash)

		} else {
			header, err := lib.HeaderByHash(c, args[0])
			if err != nil {
				return err
			}
			fmt.Printf("区块 %v 的header.ParentHash为: %v\n", header.Number, header.ParentHash)
		}

		return nil
	},
}

var tcCmd = &cobra.Command{
	Use:     "tc",
	Aliases: []string{"transactionCount"},
	Args:    cobra.ExactArgs(1),
	Short:   "统计区块包含的交易数量",
	Long:    "统计区块包含的交易数量",
	Example: "mt client block transactionCount <blockHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		count, err := lib.TransactionCount(c, args[0])
		if err != nil {
			return err
		}

		fmt.Printf("根据blockHash: %v 查询到该区块下包含 %v 笔txs。", args[0], count)
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {

			return errors.New("blockHash非法……")
		}
		return nil
	},
}
