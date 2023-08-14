// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// 测试代币，早期测试TokenLock的时候，没有用过BaybyToken，所以这里用了一个测试代币
contract Token is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 100000000 * 1e18);
    }
}