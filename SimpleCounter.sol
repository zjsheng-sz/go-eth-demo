// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title 一个简单的计数器合约
 * @dev 提供基本的计数功能：增加、减少、重置和查询
 */
contract SimpleCounter {
    // 状态变量：存储当前的计数值
    int256 private count;
    
    // 事件：当计数发生变化时触发
    event CountChanged(int256 newValue, address indexed changedBy, string operation);
    
    // 构造函数：初始化计数器（默认为0）
    constructor() {
        count = 0;
        emit CountChanged(count, msg.sender, "Initialize");
    }
    
    /**
     * @dev 将计数器加1
     */
    function increment() public {
        count += 1;
        emit CountChanged(count, msg.sender, "Increment");
    }
    
    /**
     * @dev 将计数器减1
     */
    function decrement() public {
        count -= 1;
        emit CountChanged(count, msg.sender, "Decrement");
    }
    
    /**
     * @dev 将计数器增加指定数值
     * @param _value 要增加的值（可为正数或负数）
     */
    function add(int256 _value) public {
        count += _value;
        emit CountChanged(count, msg.sender, "Add");
    }
    
    /**
     * @dev 将计数器重置为0
     */
    function reset() public {
        count = 0;
        emit CountChanged(count, msg.sender, "Reset");
    }
    
    /**
     * @dev 将计数器设置为指定值
     * @param _value 要设置的新值
     */
    function set(int256 _value) public {
        count = _value;
        emit CountChanged(count, msg.sender, "Set");
    }
    
    /**
     * @dev 获取当前计数值
     * @return 当前的计数值
     */
    function getCount() public view returns (int256) {
        return count;
    }
}