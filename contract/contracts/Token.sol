pragma solidity 0.4.24;

import "zeppelin-solidity/contracts/token/ERC20/StandardToken.sol";

contract Token is StandardToken {
    string public name = "Token";
    string public symbol = "yyy";
    uint8 public decimals = 2;
    uint256 public INITIAL_SUPPLY = 1000000000;

    constructor () public {
        totalSupply_ = INITIAL_SUPPLY;
        balances[msg.sender] = INITIAL_SUPPLY;
    }
}
