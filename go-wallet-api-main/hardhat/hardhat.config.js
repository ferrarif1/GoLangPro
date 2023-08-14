require("@nomicfoundation/hardhat-toolbox");
require("@nomiclabs/hardhat-etherscan");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    defaultNetwork: "ganache",
    networks: {
        // ganache本地测试链
        // hardhat和go测试助记词：
        // spoil test knock border hybrid drop remove hope banana face model seat
        ganache: {
            chainId: 1337,
            url: "http://127.0.0.1:8545",
            // 一些测试要用的钱包
            accounts: [
                '0x52c74ce41f0b65e5f752217ff1f1cbb58d414563b9bc4a4dcb119deb96ae36de',
                '0x407555d4ebe3a37e1906a8e610c49d92e49f108bc442f3caf8e56a5624bc3551',
            ],
            // 部署人私钥
            defaultAccount: '0x52c74ce41f0b65e5f752217ff1f1cbb58d414563b9bc4a4dcb119deb96ae36de',
        },
        // bsc测试链
        bsctestnet: {
            url: "https://data-seed-prebsc-1-s3.binance.org:8545",
            accounts: [
                '0x52c74ce41f0b65e5f752217ff1f1cbb58d414563b9bc4a4dcb119deb96ae36de',
                '0x407555d4ebe3a37e1906a8e610c49d92e49f108bc442f3caf8e56a5624bc3551',
            ],
            defaultAccount: '0x52c74ce41f0b65e5f752217ff1f1cbb58d414563b9bc4a4dcb119deb96ae36de',
        },
        // bsc主链
        bscmainnet: {
            url: "https://bsc-dataseed.binance.org/",
            accounts: [
                '0x52c74ce41f0b65e5f752217ff1f1cbb58d414563b9bc4a4dcb119deb96ae36de',
                '0x407555d4ebe3a37e1906a8e610c49d92e49f108bc442f3caf8e56a5624bc3551',
            ],
            defaultAccount: '0x52c74ce41f0b65e5f752217ff1f1cbb58d414563b9bc4a4dcb119deb96ae36de',
        }
    },
    // 编译器配置，支持同时配置多个版本
    solidity: {
        compilers: [
            {
                version: "0.8.17",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200
                    }
                }
            },
            {
                version: "0.5.16",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200
                    }
                }
            },
        ],
    },
    etherscan: {
        apiKey: "WRIA3TSVFBPXHTNHYH8D8KKX4HAFVHPDV8"
    }
};
