go-wallet-api
---

### 环境
```
nodejs18
go 1.20
[ganache](https://trufflesuite.com/ganache)
```

### ganache
```
点击NEWWORK PLACE
点击SERVER，设置：PORT NUMBER：8545，NETWORK ID：1337
点击ACCOUNTS & KEYS，设置ACCOUNT DEFAULT BALANCE：10000，覆盖默认助记词：spoil test knock border hybrid drop remove hope banana face model seat
点击右上角START即可
最小化即可，不要关闭，关闭了就是完全关闭软件了
```

### hardhat
```
cd hardhat
yarn
部署合约到ganache本地链：npx hardhat run scripts/test.js
```

### GoApi
```
cd api
# 下载依赖
go mod tidy
修改api/config/config.go中的各合约地址为上一步控制台中输出的地址
go run main.go
```

### 注意
```
部署到BSC测试网，先设置TokenLock构造参数为token地址，再：npx hardhat run scripts/deploy.js --network bsctestnet
BSC测试网每天可以领0.1个TBNB：https://testnet.bnbchain.org/faucet-smart，领完刷新页面再领，要想领更多需要VPN切换IP+多个钱包，然后将TBNB转账汇总到一起即可
生产环境需要go生成二进制包（不要基于源码运行）：go build，搭配supervisor使用，具体生成exe还是linux视情况自行百度
```

#### 接口参考API文档
#### 配置文件目录：config/config.go
#### 合约ABI转go方法，目录api/contracts/，具体见内说明
