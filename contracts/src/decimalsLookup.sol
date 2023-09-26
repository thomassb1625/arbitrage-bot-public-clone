pragma solidity 0.6.12;

abstract contract IERC20 {
    function decimals() public virtual view returns (uint8);
}

contract decimalsLookup {
    function lookupDecimals(IERC20[] calldata _tokens)
        external
        view
        returns (int256[] memory)
    {
        int256[] memory result = new int256[](_tokens.length);
        for (uint256 i = 0; i < _tokens.length; i++) {
            if (!isContract(address(_tokens[i]))) {
                result[i] = -1;
            } else {
                result[i] = _tokens[i].decimals();
            }
        }
        return result;
    }

    //Not safe when called from constructor(though not applicable to our usecase)
    function isContract(address addr) public view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(addr)
        }
        return size > 0;
    }
}
