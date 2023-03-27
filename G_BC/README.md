
- 根据rpc链接和address，返回地址是EOA还是合约


# client 基础功能
  - go run main.go client 

    打印chainID

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go client
    ChainID:  17
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c
    ChainID:  17
    ```

## client account
  - go run main.go c account isContract

    判断是否为合约地址

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a ic 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    地址: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 是否为合约: false
    ```

  - go run main.go c account balance

    查询balance

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a b 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 
    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 在最新区块的balance: 99984999999972457600528
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a b -b 0 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 在区块: 0 的balance: 100000000000000000000000
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a b -b 42 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 在区块: 42 的balance: 99984999999972457600528
    ```

## client block

  - go run main.go c block blockNum

    查询最新区块高度

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c bn
    最新区块高度为:  42
    ```

  - go run main.go c block

    查询区块详情

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b 2                                                                    
    根据num: 2 查询到了区块, block.timestamp: 1679394336
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10   
    根据blockhash: 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10 查询到了区块, block.timestamp: 1679394336
    ```

  - go run main.go c block header
  
    查询区块头信息

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b h 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10
    区块 1 的header.ParentHash为: 0xddece7f6aab5880f215b245941a5dfcd2a21d82c52d5aa356f9c0c22ee99babe
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b h 2                                                                 
    区块 2 的header.ParentHash为: 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10
    ```

  - go run main.go c block transactionCount

    查询指定区块的交易数

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b tc 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10
    根据blockHash: 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10 查询到该区块下包含 1 笔txs。
    ```

## client transaction

  - go run main.go client transaction

    查询交易信息

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c tx 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0 
    根据txHash: 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0 查询到txGas: 2000000
    ```

  - go run main.go client transaction receipt

    查询收据

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c tx r 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0
    根据txHash: 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0 查询到Receipt Status: 1
    ```

## client transfer

- go run main.go client transfer SendETHTransaction
  
  ETH transfer

  ```shell
  (anaconda3)➜  g_bc git:(main) ✗ go run main.go c t e -u https://rpc.testnet.mantle.xyz -p 7eefd641410560e690736ee331bd32512c9b58419a877eff2189facbef33cd1e
  txHash:  0xf427e40c59c93b75aa8d634086378fa7c695263eed137f00bc03a8ba326959c3
  ```

- go run main.go client transfer SendERC20Transaction

  ERC20 transfer

```shell

```

- go run main.go rpc smartContract
- go run main.go rpc eventLog
- go run main.go rpc signature
- go run main.go rpc swarm
- go run main.go rpc whisper
- go run main.go rpc utility