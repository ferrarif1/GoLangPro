package req

//WalletCreateReq 钱包创建参数
type WalletCreateReq struct {
	Count uint `form:"count,default=1" binding:"min=1,max=100"` // 创建数量
}

//WalletPrivateKeyUnlockReq 私钥解锁参数
type WalletPrivateKeyUnlockReq struct {
	PrivateKey string `form:"privateKey" binding:"required"` // 私钥
}

//WalletMnemonicUnlockReq 助记词解锁参数
type WalletMnemonicUnlockReq struct {
	Mnemonic string `form:"mnemonic" binding:"required"` // 助记词
}

// WalletGetBalanceReq 获取余额参数
type WalletGetBalanceReq struct {
	Address string `form:"address" binding:"required"` // 钱包地址
}

// WalletGetErc20BalanceReq 获取代币余额参数
type WalletGetErc20BalanceReq struct {
	Name    string `form:"name" binding:"required"`    // 代币名称
	Address string `form:"address" binding:"required"` // 钱包地址
}

// WalletTransferReq 转账参数
type WalletTransferReq struct {
	To     string  `form:"to" binding:"required"` // 收款方钱包地址
	Amount float64 `form:"amount" binding:"gt=0"` // 转账金额
}

// WalletErc20TransferReq 代币转账参数
type WalletErc20TransferReq struct {
	Name   string  `form:"name" binding:"required"` // 代币名称
	To     string  `form:"to" binding:"required"`   // 收款方钱包地址
	Amount float64 `form:"amount" binding:"gt=0"`   // 转账金额
}
