const hre = require("hardhat");
const ethers = hre.ethers;

// 用于测试API的本地部署脚本
async function main() {
    const accounts = await ethers.getSigners();

    console.log("\n", "\n");

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
    console.log("1.1、设置NFT合约的最大供应量为70000");

    // 铸造NFT 1给accounts[1] (首次tokenId默认为1，第二个参数代表数量)
    await nft.mintSeaDrop(
        accounts[1].address, 
        1
    );
    console.log("1.2、铸造NFT 1 给accounts[1]");

    // 验证NFT 1的拥有者是否为accounts[1]
    const owner = await nft.ownerOf(1);
    console.log("1.3、验证NFT 1的拥有者是否为accounts[1]: ", owner == accounts[1].address);

    // account[1]将NFT授权给account[0]，方便测试授权转账API
    // 非部署人调用合约要像这样
    await new ethers.Contract(nft.address, nft.interface, accounts[1]).approve(
        accounts[0].address, 
        1
    );
    console.log("1.4、account[1]将NFT 1授权给account[0]，方便测试授权转账API", "\n", "\n");
    //---------------------------------NFT---------------------------------


    //---------------------------------USDT---------------------------------
    // 部署USDT合约
    const BEP20USDT = await ethers.getContractFactory("BEP20USDT");
    const usdt = await BEP20USDT.deploy();
    console.log("2、USDT合约部署成功: ", usdt.address, "\n", "\n");
    //---------------------------------USDT---------------------------------


    //---------------------------------测试Token1---------------------------
    // 部署Token合约
    const Token = await ethers.getContractFactory("Token");
    const token = await Token.deploy(
        "AAA",
        "AAA",
    );
    console.log("3、Token合约部署成功: ", token.address, "\n", "\n");
    //---------------------------------测试Token1---------------------------------

    //---------------------------------测试Token2---------------------------
    // 部署Token2合约，合约与Token1不同，证明API接口可以ERC20标准代币通用
    const Token2 = await ethers.getContractFactory("Token2");
    const token2 = await Token2.deploy(
        "BBB",
        "BBB",
    );
    console.log("4、Token2合约部署成功: ", token2.address, "\n", "\n");
    //---------------------------------测试Token2---------------------------------


    //---------------------------------TokenLock---------------------------------
    // 部署TokenLock合约
    const TokenLock = await ethers.getContractFactory("TokenLock");
    // 传入Token合约地址
    const tokenLock = await TokenLock.deploy(token.address);
    console.log("5、TokenLock合约部署成功: ", tokenLock.address);

    // 部署人转账给TokenLock合约1000000个Token用于锁仓
    console.log("5.1、部署人转账1000000个Token给TokenLock合约，模拟待解锁总额，方便测试解锁API");
    await token.transfer(
        tokenLock.address, ethers.utils.parseEther("1000000"),
    );

    // 验证TokenLock合约的余额
    const tokenLockBalance = await token.balanceOf(tokenLock.address);
    console.log("5.2、验证TokenLock合约的余额: ", ethers.utils.formatEther(tokenLockBalance), "\n", "\n");
    //---------------------------------TokenLock---------------------------------
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });