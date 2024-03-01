pragma solidity >=0.4.25 <0.6.0;
// over 8.0 , eth will call some method only for EVM???

import "./ConvertLib.sol";

contract MetaCoin {
    mapping (address => uint) balances;
    mapping (string => int256) myMapping;

    function hi() public returns (bool sufficient) {
        myMapping["hello"] = 22222222;
        return true;
    }

    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    constructor() public {
        balances[tx.origin] = 10000;
        myMapping["hello"] = 1111;
    }

    function sendCoin(address receiver, uint amount) public returns(bool sufficient) {
        if (balances[msg.sender] < amount) return false;
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        emit Transfer(msg.sender, receiver, amount);
        return true;
    }

    function getBalanceInEth(address addr) public view returns(uint){
        return ConvertLib.convert(getBalance(addr),2);
    }

    function getBalance(address addr) public view returns(uint) {
        return balances[addr];
    }
}
