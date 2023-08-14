package logic

import (
	"math/big"
	"server/config"
	"server/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	NftLogic    = nftLogic{}
	nftInstance = CommonLogic.GetNftInstance()
)

// nftLogic NFT逻辑层
type nftLogic struct{}

// MintSeaDrop 铸造NFT
func (nl nftLogic) MintSeaDrop(minter string, quantity uint) (txHash string) {
	// 铸造人私钥解锁，获取钱包
	fromAddress := WalletLogic.PrivateKeyUnlock(config.Config.NftFromPrivateKey)
	// 校验验证接收人钱包格式
	CommonLogic.CheckAddress(minter, "接收人")
	// 创建签名选项
	auth := CommonLogic.GetAuth(config.Config.NftFromPrivateKey, fromAddress)
	// 生成未签名事务
	tx, err := nftInstance.MintSeaDrop(
		auth,
		common.HexToAddress(minter),
		big.NewInt(int64(quantity)),
	)
	util.CheckUtil.CheckApiErr(err, "铸造NFT失败")
	// 返回txHash
	txHash = tx.Hash().Hex()
	return
}

// GetApproved 获取NFT授权地址
func (nl nftLogic) GetApproved(tokenId uint) (address string) {
	// 获取授权地址
	approved, err := nftInstance.GetApproved(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	util.CheckUtil.CheckApiErr(err, "获取NFT授权地址失败")
	address = approved.Hex()
	return
}

// GetOwnerOf 获取NFT持有人
func (nl nftLogic) GetOwnerOf(tokenId uint) (address string) {
	// 获取授权地址
	approved, err := nftInstance.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	util.CheckUtil.CheckApiErr(err, "获取NFT持有人失败")
	address = approved.Hex()
	return
}

// TransferFrom NFT转账
func (nl nftLogic) TransferFrom(from string, to string, tokenId uint) (txHash string) {
	// 调用人私钥解锁，获取钱包
	toAddress := WalletLogic.PrivateKeyUnlock(config.Config.NftFromPrivateKey)
	// 校验发送人
	CommonLogic.CheckAddress(from, "发送人")
	// 校验接收人
	CommonLogic.CheckAddress(to, "接收人")
	// 创建签名选项
	auth := CommonLogic.GetAuth(config.Config.NftFromPrivateKey, toAddress)
	// 发起转账
	tx, err := nftInstance.TransferFrom(
		auth,
		common.HexToAddress(from),
		common.HexToAddress(to),
		big.NewInt(int64(tokenId)),
	)
	util.CheckUtil.CheckApiErr(err, "NFT转账失败")
	// 返回txHash
	txHash = tx.Hash().Hex()
	return
}
