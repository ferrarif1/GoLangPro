package go_cowswap

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	contract_binding "github.com/itsahedge/go-cowswap/contracts/generated"
	"math/big"
	"strconv"
)

// GetAllowance - Token allowance for owner
func (c *Client) GetAllowance(ctx context.Context, ownerAddress, tokenAddress string) (*big.Int, error) {
	if ownerAddress == "" {
		return nil, errors.New("must provide an Owner Address")
	}
	if tokenAddress == "" {
		return nil, errors.New("must provide a Token Address")
	}
	token := common.HexToAddress(tokenAddress)
	contract, err := contract_binding.NewErc20(token, c.EthClient)
	if err != nil {
		return nil, err
	}
	owner := common.HexToAddress(ownerAddress)
	spender := common.HexToAddress(GPv2_Vault_Relayer)
	allowance, err := contract.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
	if err != nil {
		return nil, err
	}
	return allowance, nil
}

type ApproveAllowance struct {
	TokenAddress  string `json:"token_address"`
	WalletAddress string `json:"wallet_address"`
}

// SetAllowance - Set token allowance for owner on CowSwap GPv2 Vault Relayer contract
func (c *Client) SetAllowance(ctx context.Context, tokenAddress, tokenAmount string) (*types.Transaction, error) {
	if c.TransactionSigner == nil {
		return nil, errors.New("transaction signer not initialized")
	}
	var amountToApprove *big.Int
	if tokenAmount == "" {
		amountToApprove = new(big.Int).Lsh(big.NewInt(1), 256-7)
	}
	if tokenAmount != "" {
		i, _ := strconv.ParseInt(tokenAmount, 10, 64)
		amountToApprove = new(big.Int).Set(big.NewInt(i))
	}
	auth := c.TransactionSigner.Auth
	token := common.HexToAddress(tokenAddress)
	contract, err := contract_binding.NewErc20(token, c.EthClient)
	if err != nil {
		panic(err)
	}
	spender := common.HexToAddress(GPv2_Vault_Relayer)
	opts := &bind.TransactOpts{
		Context: ctx,
		Signer:  auth.Signer,
		From:    auth.From,
	}
	tx, err := contract.Approve(opts, spender, amountToApprove)
	if err != nil {
		return nil, fmt.Errorf("Approve() err:\n%v", err)
	}
	return tx, nil
}
