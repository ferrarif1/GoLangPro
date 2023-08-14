const { expect } = require("chai");

describe("TokenLock", function () {
    let token;
    let tokenLock;
    let accounts;

    const eth90000000 = ethers.utils.parseEther("90000000");
    const eth21600 = ethers.utils.parseEther("21600");
    const eth20000 = ethers.utils.parseEther("20000");
    const eth1600 = ethers.utils.parseEther("1600");

    before(async function () {
        const Token = await ethers.getContractFactory("Token");
        const TokenLock = await ethers.getContractFactory("TokenLock");

        token = await Token.deploy("AAA", "AAA");
        await token.deployed();

        tokenLock = await TokenLock.deploy(token.address);
        await tokenLock.deployed();
        
        token.transfer(tokenLock.address, eth90000000);
        accounts = await ethers.getSigners();
    });

    it("检查TokenLock持有代币是否为90000000", async function () {
        const balance = await token.balanceOf(tokenLock.address);
        expect(balance).to.equal(eth90000000);
    });

    it("触发今日解锁，检查已解锁代币是否为21600", async function () {
        await tokenLock.unlock({ from: accounts[0].address });
        const unlockedToken = await tokenLock.unlockedToken();
        expect(unlockedToken).to.equal(eth21600);
    });

    it("从合约转账20000个已解锁代币给account[1]，检查剩余已解锁代币是否为1600", async function () {
        await tokenLock.transfer(
            accounts[1].address,
            eth20000,
            { from: accounts[0].address }
        );

        const unlockedToken = await tokenLock.unlockedToken();
        expect(unlockedToken).to.equal(eth1600);

        const balance = await token.balanceOf(accounts[1].address);
        expect(balance).to.equal(eth20000);
    });

    it("检查account[1]余额是否为20000", async function () {
        const balance = await token.balanceOf(accounts[1].address);
        expect(balance).to.equal(eth20000);
    });
});