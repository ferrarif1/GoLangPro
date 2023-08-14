const hre = require("hardhat");
const ethers = hre.ethers;

// 正式部署上链脚本demo
// 部署方法：npx hardhat run scripts/deploy.js --network bsc_testnet
// 要求：部署前先在hardhat.config.js中配置好defaultAccount，这是合约部署人
async function main() {
    const accounts = await ethers.getSigners();

    console.log("\n");

    //---------------------------------NFT---------------------------------
    // 部署NFT合约，设置NFT name、symbol、拥有铸造权的地址组
    const ERC721SeaDrop = await ethers.getContractFactory("ERC721SeaDrop");
    const nft = await ERC721SeaDrop.deploy(
        "MyNft", 
        "MNFT", 
        [accounts[0].address]   //[铸造人1,铸造人2]
    );
    console.log("1、NFT合约部署成功: ", nft.address);

    // 设置NFT最大供应量，不设置无法铸造，切记
    await nft.setMaxSupply(70000);
    console.log("1.1、设置NFT合约的最大供应量为70000", "\n", "\n");
    //---------------------------------NFT---------------------------------

    //---------------------------------TokenLock---------------------------------
    // 部署TokenLock合约，需要填入你的线上正式baytoken地址
    const TokenLock = await ethers.getContractFactory("TokenLock");
    // 传入Token合约地址，正式使用请传具体token合约地址
    const tokenLock = await TokenLock.deploy("0x55d398326f99059fF775485246999027B3197955");
    console.log("2、TokenLock合约部署成功: ", tokenLock.address);

    // 部署人转账给TokenLock合约1000000个（假设是这么多，具体自己改）Token用于锁仓
    console.log("2.1、部署人转账1000000个Token给TokenLock合约，表示待解锁总额");
    await token.transfer(
        tokenLock.address, 
        ethers.utils.parseEther("1000000"),
    );

    // 验证TokenLock合约的余额
    const tokenLockBalance = await token.balanceOf(tokenLock.address);
    console.log("2.2、验证TokenLock合约的余额: ", ethers.utils.formatEther(tokenLockBalance), "\n", "\n");
    // ---------------------------------TokenLock---------------------------------
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });