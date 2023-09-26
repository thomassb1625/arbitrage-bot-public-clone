//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;
pragma experimental ABIEncoderV2;

import {IPoolAddressesProvider} from "@aave-v3-core/contracts/interfaces/IPoolAddressesProvider.sol";
import {IERC20} from "@aave-v3-core/contracts/dependencies/openzeppelin/contracts/IERC20.sol";

interface IGLMR is IERC20 {
    function deposit() external payable;
    function withdraw(uint256) external;
}

contract Executor {

    // constructor(address _addressProvider)
    //     FlashLoanSimpleReceiverBase(IPoolAddressesProvider(_addressProvider))
    // {
    //     owner = payable(msg.sender);
    // }
    address payable public immutable owner;
    address public immutable executor;
    IGLMR private immutable WGLMR;
    bool private flashLoanGuard;

    constructor(address _executor, address _wglmr) payable {
        WGLMR = IGLMR(_wglmr);
        owner = payable(msg.sender);
        executor = _executor;
        if (msg.value > 0) {
            WGLMR.deposit{value: msg.value}();
        }
    }


    //Same as old logic
    function swapWithoutFlashLoan(
        uint256 _wglmrAmountToFirstMarket,
        uint256 _glmrAmountToCoinbase,
        address[] memory _targets,
        bytes[] memory _payloads
    ) external payable onlyExecutor {
        require(_targets.length == _payloads.length, "WL");
        //Prevents flashloan from being called
        flashLoanGuard = true;
        uint256 _wglmrBalanceBefore = WGLMR.balanceOf(address(this));
        WGLMR.transfer(_targets[0], _wglmrAmountToFirstMarket);
        for (uint256 i = 0; i < _targets.length; i++) {
            (bool _success, bytes memory _response) = _targets[i].call(
                _payloads[i]
            );
            require(_success, "CE");
            _response;
        }

        uint256 _wglmrBalanceAfter = WGLMR.balanceOf(address(this));
        require(
            _wglmrBalanceAfter > _wglmrBalanceBefore + _glmrAmountToCoinbase,
            "B0"
        );
        if (_glmrAmountToCoinbase == 0) return;

        uint256 _glmrBalance = address(this).balance;
        if (_glmrBalance < _glmrAmountToCoinbase) {
            WGLMR.withdraw(_glmrAmountToCoinbase - _glmrBalance);
        }
        block.coinbase.transfer(_glmrAmountToCoinbase);
        flashLoanGuard = false;
    }

    function getBalance(address _tokenAddress) external view returns (uint256) {
        return IERC20(_tokenAddress).balanceOf(address(this));
    }

    function withdraw(address _tokenAddress) external onlyOwner {
        IERC20 token = IERC20(_tokenAddress);
        token.transfer(msg.sender, token.balanceOf(address(this)));
    }

    function withdrawGLMR() external onlyOwner {
        owner.transfer(address(this).balance);
    }

    modifier onlyExecutor() {
        require(msg.sender == executor, "ONee");
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    receive() external payable {}
}

