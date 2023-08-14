// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract TokenLock {
    // 代币实例
    IERC20 public token;
    // 代币精度18（10的18次方）
    uint256 private constant DECIMAL_MULTIPLIER = 1e18;
    // 初始每天解锁数量
    uint256 private constant INITIAL_DAILY_UNLOCK = 21600 * DECIMAL_MULTIPLIER;
    // 减半周期90天
    uint256 private constant PERIOD_DAYS = 90;
    // 360天后不再减半
    uint256 private constant TOTAL_DAYS = 360;
    // 360天后每天固定解锁1000个
    uint256 private constant MIN_DAILY_UNLOCK = 1000 * DECIMAL_MULTIPLIER;

    // 解锁启动时间
    uint256 public startTimestamp;
    // 已解锁的数量
    uint256 public unlockedToken;
    // 上一次解锁时间
    uint256 public lastUnlockTimestamp;
    // 合约所有者
    address public owner;

    // 修饰器，只有合约所有者才能调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }

    // 解锁事件
    event UnlockEvent(uint256 indexed amount, uint256 indexed timestamp);
    // 转账事件
    event TransferEvent(address indexed to, uint256 indexed amount);

    // 构造函数
    // _token: 代币地址
    constructor(address _token) {
        // 代币地址不能是零地址
        require(_token != address(0), "Token address cannot be zero address.");
        // 实例化代币合约
        token = IERC20(_token);

        // 需要减一天，部署完可以立马解锁一次
        // 否则因为unlock()中不满1天直接return了
        startTimestamp = block.timestamp - 1 days;

        lastUnlockTimestamp = startTimestamp;
        // 设置合约所有者
        owner = msg.sender;
    }

    // 代币解锁
    // 定时任务触发，多次触发无影响
    function unlock() external onlyOwner {
        // 距离上次解锁时间必须大于1天
        require(block.timestamp - lastUnlockTimestamp >= 1 days, "The time since the last unlock is less than one day.");
        // 合约代币余额至少大于1000
        require(token.balanceOf(address(this)) >= MIN_DAILY_UNLOCK, "Insufficient balance");
        // 计算自启动以来过去了多少天
        uint256 daysSinceStart = (block.timestamp - startTimestamp) / 1 days;
        // 计算自启动以来减半了多少次
        uint256 periodsSinceStart = daysSinceStart / PERIOD_DAYS;
        // 计算当前应解锁的数量
        uint256 dailyUnlock;

        // 如果自启动以来过去了超过360天，则每天固定解锁1000个token
        if (daysSinceStart >= TOTAL_DAYS) {
            dailyUnlock = MIN_DAILY_UNLOCK;
        } else {
            // 否则，每天解锁的数量随着时间的推移而减半
            dailyUnlock = INITIAL_DAILY_UNLOCK / 2 ** periodsSinceStart;
        }

        // 计算自上一次解锁以来过去了多少天
        uint256 daysSinceLastUnlock = daysSinceStart - (lastUnlockTimestamp - startTimestamp) / 1 days;
        // 计算当前待解锁的代币数量
        uint256 newUnlockedToken = daysSinceLastUnlock * dailyUnlock;

        // 已解锁的数量增加
        unlockedToken += newUnlockedToken;
        // 记录本次解锁时间
        lastUnlockTimestamp = block.timestamp;

        // 触发解锁事件
        emit UnlockEvent(newUnlockedToken, block.timestamp);
    }

    // 将已解锁的代币转账给指定地址
    function transfer(address to, uint256 amount) external onlyOwner {
        // 接收人不能时零地址和本合约地址
        require(to != address(0), "Recipient address cannot be zero address.");
        require(to != address(this), "Recipient address cannot be contract address.");
        // 已解锁数量必须 >= 转账数量
        require(unlockedToken >= amount, "Not enough unlocked tokens.");

        // 扣除已解锁数量
        unlockedToken -= amount;
        // 转账
        token.transfer(to, amount);

        // 触发转账事件
        emit TransferEvent(to, amount);
    }
}