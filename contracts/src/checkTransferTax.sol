//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.6.12;

pragma experimental ABIEncoderV2;

import {IERC20} from "protocol-v2/contracts/dependencies/openzeppelin/contracts/IERC20.sol";

contract checkTax {
    // _wethAmountToFirstMarket is no longer used
    function swapWithoutFlashLoan(
        address[] memory _targets,
        bytes[] memory _payloads,
        address token
    ) external payable returns (uint256) {
        for (uint256 i = 0; i < _targets.length; i++) {
            //Don't assign values, saves a bit of gas
            (bool _success, bytes memory _response) = _targets[i].call(_payloads[i]);
            require(_success, "Failed here");
            //_response;
        }
        return IERC20(token).balanceOf(address(this));
    }
}
