package resp

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// WalletCreateResp 钱包创建返回信息
type WalletCreateResp struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	Mnemonic   string `json:"mnemonic"`
}

// WalletPrivateKeyUnlockResp 私钥解锁返回信息
type WalletPrivateKeyUnlockResp struct {
	Address string `json:"address"`
}

// WalletMnemonicUnlockResp 助记词解锁返回信息
type WalletMnemonicUnlockResp struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

// WalletBalanceResp 余额返回信息
type WalletBalanceResp struct {
	Balance decimal.Decimal `json:"balance"`
}

// WalletTxResp 返回Tx
type WalletTxResp struct {
	Tx string `json:"tx"`
}

// WalletNumResp 已解锁代币数量返回信息
type WalletNumResp struct {
	Num decimal.Decimal `json:"num"`
}

// WalletTimestampResp 时间戳返回信息
type WalletTimestampResp struct {
	Timestamp *big.Int `json:"timestamp"`
}
