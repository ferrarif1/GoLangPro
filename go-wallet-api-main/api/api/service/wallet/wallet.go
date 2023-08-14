package wallet

import (
	"server/api/logic"
	"server/api/schemas/resp"
)

var WalletService = walletService{}

// walletService 钱包服务实现类
type walletService struct{}

// Create 生成钱包
// Count: 生成数量，不填默认为1，最大100 (验证器中有限制)
func (ws walletService) Create(Count uint) (data []resp.WalletCreateResp) {
	data = logic.WalletLogic.BatchCreateWallet(Count)
	return
}

// PrivateKeyUnlock 私钥解锁钱包
func (ws walletService) PrivateKeyUnlock(privateKey string) resp.WalletPrivateKeyUnlockResp {
	address := logic.WalletLogic.PrivateKeyUnlock(privateKey)
	return resp.WalletPrivateKeyUnlockResp{
		Address: address,
	}
}

// MnemonicUnlock 助记词解锁钱包
func (ws walletService) MnemonicUnlock(mnemonic string) resp.WalletMnemonicUnlockResp {
	address, privateKey := logic.WalletLogic.MnemonicUnlock(mnemonic)
	return resp.WalletMnemonicUnlockResp{
		Address:    address,
		PrivateKey: privateKey,
	}
}

// GetBalance 查询余额
func (ws walletService) GetBalance(address string) resp.WalletBalanceResp {
	balance := logic.CommonLogic.GetBalance(address)
	return resp.WalletBalanceResp{
		Balance: balance,
	}
}

// GetErc20Balance 查询代币余额
func (ws walletService) GetErc20Balance(name string, address string) resp.WalletBalanceResp {
	balance := logic.CommonLogic.GetErc20Balance(name, address)
	return resp.WalletBalanceResp{
		Balance: balance,
	}
}

// Transfer 主网币转账
func (ws walletService) Transfer(to string, amount float64) resp.WalletTxResp {
	tx := logic.WalletLogic.Transfer(to, amount)
	return resp.WalletTxResp{
		Tx: tx,
	}
}

// Erc20Transfer 代币转账
func (ws walletService) Erc20Transfer(name string, to string, amount float64) resp.WalletTxResp {
	tx := logic.WalletLogic.Erc20Transfer(name, to, amount)
	return resp.WalletTxResp{
		Tx: tx,
	}
}

// Unlock TokenLock解锁合约代币
func (ws walletService) Unlock() resp.WalletTxResp {
	tx := logic.WalletLogic.Unlock()
	return resp.WalletTxResp{
		Tx: tx,
	}
}

// GetUnlockToken 获取TokenLock已解锁代币数量
func (ws walletService) GetUnlockToken() resp.WalletNumResp {
	num := logic.WalletLogic.GetUnlockToken()
	return resp.WalletNumResp{
		Num: num,
	}
}

// GetLastUnlockTimestamp 获取上次解锁时间戳
func (ws walletService) GetLastUnlockTimestamp() resp.WalletTimestampResp {
	timestamp := logic.WalletLogic.GetLastUnlockTimestamp()
	return resp.WalletTimestampResp{
		Timestamp: timestamp,
	}
}

// GetStartTimestamp 获取合约启动时间戳
func (ws walletService) GetStartTimestamp() resp.WalletTimestampResp {
	startTimestamp := logic.WalletLogic.GetStartTimestamp()
	return resp.WalletTimestampResp{
		Timestamp: startTimestamp,
	}
}

// TokenLockTransfer 已解锁代币转账
func (ws walletService) TokenLockTransfer(to string, amount float64) resp.WalletTxResp {
	tx := logic.WalletLogic.TokenLockTransfer(to, amount)
	return resp.WalletTxResp{
		Tx: tx,
	}
}
