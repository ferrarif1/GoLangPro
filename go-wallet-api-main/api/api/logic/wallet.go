package logic

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"server/api/schemas/resp"
	"server/config"
	"server/util"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/shopspring/decimal"
	bip39 "github.com/tyler-smith/go-bip39"
)

var (
	WalletLogic       = walletLogic{}
	tokenLockInstance = CommonLogic.GetTokenLockInstance()
)

// walletLogic 钱包逻辑层
type walletLogic struct{}

// CreateWallet 单个钱包生成
func (wl walletLogic) CreateWallet() (address, privateKey, mnemonic string) {
	// 创建128位随机熵字节
	entropy, err := bip39.NewEntropy(128)
	util.CheckUtil.CheckApiErr(err, "创建随机熵字节失败")
	// 生成助记词
	mnemonic, _ = bip39.NewMnemonic(entropy)
	// 通过助记词生成钱包地址、私钥
	address, privateKey = wl.MnemonicUnlock(mnemonic)
	return
}

// BatchCreateWallet 批量生成钱包
func (wl walletLogic) BatchCreateWallet(Count uint) (data []resp.WalletCreateResp) {
	count := int(Count)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(count)
	for i := 0; i < count; i++ {
		// 协程处理
		go func() {
			defer wg.Done()
			address, privateKey, mnemonic := wl.CreateWallet()
			mu.Lock()
			data = append(data, resp.WalletCreateResp{
				Address:    address,
				PrivateKey: privateKey,
				Mnemonic:   mnemonic,
			})
			mu.Unlock()
		}()
	}
	wg.Wait()
	return
}

// PrivateKeyUnlock 私钥解锁钱包（得到钱包地址）
func (wl walletLogic) PrivateKeyUnlock(privKey string) (address string) {
	// 如果私钥以0x开头则去掉，否则原样返回
	privKey = strings.TrimPrefix(privKey, "0x")
	privateKey := CommonLogic.GetTruePrivateKey(privKey)
	publicKey := privateKey.Public()
	if publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey); ok {
		address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	}
	return
}

// MnemonicUnlock 助记词解锁钱包（得到钱包地址+私钥）
func (wl walletLogic) MnemonicUnlock(mnemonic string) (address string, privateKey string) {
	if ok := bip39.IsMnemonicValid(mnemonic); !ok {
		util.CheckUtil.CheckApiErr(errors.New(""), "助记词无效")
	}
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := hdwallet.NewFromSeed(seed)
	util.CheckUtil.CheckApiErr(err, "从BIP-39种子生成钱包失败")
	// 设置标准派生路径
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	util.CheckUtil.CheckApiErr(err, "派生新钱包失败")
	// 生成钱包地址和私钥
	address = account.Address.Hex()
	privateKey, _ = wallet.PrivateKeyHex(account)
	return
}

// Transfer 主网币转账
func (wl walletLogic) Transfer(to string, amount float64) (txHash string) {
	// 发款方私钥解锁，获取钱包地址
	fromAddress := wl.PrivateKeyUnlock(config.Config.PrivateKey)
	// 校验钱包余额
	CommonLogic.CheckBalance(fromAddress, amount)
	// 校验收款方钱包格式
	CommonLogic.CheckAddress(to, "收款方")
	// 生成未签名事务
	tx := CommonLogic.GetNewTransaction(fromAddress, to, amount, nil)
	// 私钥签名
	signedTx := CommonLogic.GetSign(tx, config.Config.PrivateKey)
	// 广播交易
	CommonLogic.Send(signedTx)
	// 返回txHash
	txHash = signedTx.Hash().Hex()
	return
}

// Erc20Transfer 代币转账
func (wl walletLogic) Erc20Transfer(name string, to string, amount float64) (txHash string) {
	// 获取代币信息
	info := CommonLogic.GetErc20Info(name)
	// 获取代币实例
	erc20Instance := CommonLogic.GetErc20Instance(name)
	// 发款方私钥解锁，获取钱包地址
	fromAddress := wl.PrivateKeyUnlock(info.PrivateKey)
	// 校验发送人余额
	CommonLogic.CheckErc20Balance(name, fromAddress, amount)
	// 校验收款方钱包格式
	CommonLogic.CheckAddress(to, "收款方")
	// 创建签名选项
	auth := CommonLogic.GetAuth(info.PrivateKey, fromAddress)
	// 转账
	tx, err := erc20Instance.Transfer(
		auth,
		common.HexToAddress(to),
		CommonLogic.ToWei(amount, int(info.Decimal)),
	)
	util.CheckUtil.CheckApiErr(err, "转账失败")
	// 返回txHash
	txHash = tx.Hash().Hex()
	return
}

// Unlock 解锁TokenLock合约代币
func (wl walletLogic) Unlock() (txHash string) {
	// 合约所有人私钥解锁，获取钱包地址
	ownerAddress := wl.PrivateKeyUnlock(config.Config.TokenLockPrivateKey)
	// 创建签名选项
	auth := CommonLogic.GetAuth(config.Config.TokenLockPrivateKey, ownerAddress)
	// 解锁
	tx, err := tokenLockInstance.Unlock(auth)
	util.CheckUtil.CheckApiErr(err, "解锁失败")
	// 返回txHash
	txHash = tx.Hash().Hex()
	return
}

// GetUnlockToken 获取已解锁的代币数量
func (wl walletLogic) GetUnlockToken() (result decimal.Decimal) {
	unLockToken, err := tokenLockInstance.UnlockedToken(&bind.CallOpts{})
	util.CheckUtil.CheckApiErr(err, "获取已解锁代币数量失败")
	result = CommonLogic.ToDecimal(unLockToken, int(config.Config.TokenLockDecimal))
	return
}

// GetLastUnlockTimestamp 获取上次解锁时间戳
func (wl walletLogic) GetLastUnlockTimestamp() (timestamp *big.Int) {
	timestamp, err := tokenLockInstance.LastUnlockTimestamp(&bind.CallOpts{})
	util.CheckUtil.CheckApiErr(err, "获取上次解锁时间戳失败")
	return
}

// GetStartTimestamp 获取合约启动时间戳
func (wl walletLogic) GetStartTimestamp() (startTimestamp *big.Int) {
	startTimestamp, err := tokenLockInstance.StartTimestamp(&bind.CallOpts{})
	util.CheckUtil.CheckApiErr(err, "获取解锁启动时间戳失败")
	return
}

// TokenLockTransfer 已解锁代币转账
func (wl walletLogic) TokenLockTransfer(to string, amount float64) (txHash string) {
	// 部署人私钥解锁，获取钱包地址
	ownerAddress := wl.PrivateKeyUnlock(config.Config.TokenLockPrivateKey)
	// 校验收款方钱包格式
	CommonLogic.CheckAddress(to, "收款方")
	// 创建签名选项
	auth := CommonLogic.GetAuth(config.Config.TokenLockPrivateKey, ownerAddress)
	// 转账
	tx, err := tokenLockInstance.Transfer(
		auth,
		common.HexToAddress(to),
		CommonLogic.ToWei(amount, int(config.Config.TokenLockDecimal)),
	)
	util.CheckUtil.CheckApiErr(err, "转账失败")
	// 返回txHash
	txHash = tx.Hash().Hex()
	return
}
