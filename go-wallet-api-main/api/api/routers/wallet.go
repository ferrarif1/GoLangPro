package routers

import (
	"server/api/schemas/req"
	"server/api/service/wallet"
	"server/core"
	"server/core/response"
	"server/util"

	"github.com/gin-gonic/gin"
)

var WalletGroup = core.Group("/wallet")

func init() {
	group := WalletGroup
	group.AddGET("/create", create)
	group.AddPOST("/privateKeyUnlock", privateKeyUnlock)
	group.AddPOST("/mnemonicUnlock", mnemonicUnlock)

	group.AddPOST("/getBalance", getBalance)
	group.AddPOST("/getErc20Balance", getErc20Balance)

	group.AddPOST("/transfer", transfer)
	group.AddPOST("/erc20Transfer", erc20Transfer)

	group.AddGET("/unlock", unlock)
	group.AddGET("/getUnlockToken", getUnlockToken)
	group.AddGET("/getLastUnlockTimestamp", getLastUnlockTimestamp)
	group.AddGET("/getStartTimestamp", getStartTimestamp)
	group.AddPOST("/tokenLockTransfer", tokenLockTransfer)
}

// 生成钱包
func create(c *gin.Context) {
	var createReq req.WalletCreateReq
	util.VerifyUtil.VerifyQuery(c, &createReq)
	response.OkWithData(c, wallet.WalletService.Create(createReq.Count))
}

// 私钥解锁
func privateKeyUnlock(c *gin.Context) {
	var privateKeyReq req.WalletPrivateKeyUnlockReq
	util.VerifyUtil.VerifyBody(c, &privateKeyReq)
	response.OkWithData(c, wallet.WalletService.PrivateKeyUnlock(privateKeyReq.PrivateKey))
}

// 助记词解锁
func mnemonicUnlock(c *gin.Context) {
	var mnemonicReq req.WalletMnemonicUnlockReq
	util.VerifyUtil.VerifyBody(c, &mnemonicReq)
	response.OkWithData(c, wallet.WalletService.MnemonicUnlock(mnemonicReq.Mnemonic))
}

// 查询主网币余额
func getBalance(c *gin.Context) {
	var getBalanceReq req.WalletGetBalanceReq
	util.VerifyUtil.VerifyBody(c, &getBalanceReq)
	response.OkWithData(c, wallet.WalletService.GetBalance(getBalanceReq.Address))
}

// 查询代币余额
func getErc20Balance(c *gin.Context) {
	var getErc20BalanceReq req.WalletGetErc20BalanceReq
	util.VerifyUtil.VerifyBody(c, &getErc20BalanceReq)
	response.OkWithData(c, wallet.WalletService.GetErc20Balance(getErc20BalanceReq.Name, getErc20BalanceReq.Address))
}

// 主网币转账(ETH/BNB等)
func transfer(c *gin.Context) {
	var transferReq req.WalletTransferReq
	util.VerifyUtil.VerifyBody(c, &transferReq)
	response.OkWithData(c, wallet.WalletService.Transfer(transferReq.To, transferReq.Amount))
}

// 代币转账
func erc20Transfer(c *gin.Context) {
	var erc20TransferReq req.WalletErc20TransferReq
	util.VerifyUtil.VerifyBody(c, &erc20TransferReq)
	response.OkWithData(c, wallet.WalletService.Erc20Transfer(erc20TransferReq.Name, erc20TransferReq.To, erc20TransferReq.Amount))
}

// TokenLock合约代币解锁
func unlock(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.Unlock())
}

// 获取TokenLock合约已解锁代币数量
func getUnlockToken(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.GetUnlockToken())
}

// 获取TokenLock上次解锁时间
func getLastUnlockTimestamp(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.GetLastUnlockTimestamp())
}

// 获取TokenLock合约启动时间
func getStartTimestamp(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.GetStartTimestamp())
}

// 已解锁代币转账
func tokenLockTransfer(c *gin.Context) {
	var tokenLockTransferReq req.WalletTransferReq
	util.VerifyUtil.VerifyBody(c, &tokenLockTransferReq)
	response.OkWithData(c, wallet.WalletService.TokenLockTransfer(tokenLockTransferReq.To, tokenLockTransferReq.Amount))
}
