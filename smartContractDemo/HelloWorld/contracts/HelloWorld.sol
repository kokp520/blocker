// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract HelloWorld {
    mapping (string => int256) test;

    constructor(){
        test["world"] = 99123;
    }

    function hi() public returns (bool sufficient) {
        test["world"] = 123;
        return true;
    }
}
