pragma solidity ^0.4.24;

import "./ERC20Basic.sol";

/// @title A title that should describe the contract/interface
/// @author The name of the author
/// @notice Explain to an end user what this does
/// @dev Explain to a developer any extra details

contract ERC20 is ERC20Basic {
    function allowance(address _owner, address _spender) public view returns (uint256);

    function transferFrom(address _from, address _to, uint256 _value) public returns(bool);

    function approve(address _spender, uint256 _value) public returns (bool);

    event Approve(
        address indexed owner,
        address indexed spender,
        uint256 value
    )   
}