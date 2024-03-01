pragma solidity ^0.4.24;

/// @title A title that should describe the contract/interface
/// @author The name of the author
/// @notice Explain to an end user what this does
/// @dev Explain to a developer any extra detailsa


contract ERC20Basic 
{
    // get total coin method.
    function totalSupply() public view returns (uint256); 

    // get address balance method.
    function balanceOf(address _who) public view returns (uint256);

    //MUST check address coin first 
    function transfer(address _to, uint256 _value) public returns (bool);

    // record tx
    event Transfer(
        address indexed from,
        address indexed to,
        uint256 value
    );
}