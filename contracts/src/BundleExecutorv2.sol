//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.6.12;
pragma experimental ABIEncoderV2;

import {FlashLoanReceiverBase} from "protocol-v2/contracts/flashloan/base/FlashLoanReceiverBase.sol";
import {ILendingPoolAddressesProvider} from "protocol-v2/contracts/interfaces/ILendingPoolAddressesProvider.sol";
import {IERC20} from "protocol-v2/contracts/dependencies/openzeppelin/contracts/IERC20.sol";

interface IWETH is IERC20 {
    function deposit() external payable;

    function withdraw(uint256) external;
}

contract BundleExecutorV2 is FlashLoanReceiverBase {
    address payable public immutable owner;
    address public immutable executor;
    IWETH private WETH;
    bool private flashLoanGuard;

    constructor(
        address _executor,
        address _weth,
        address _provider
    )
        public
        payable
        FlashLoanReceiverBase(ILendingPoolAddressesProvider(_provider))
    {
        WETH = IWETH(_weth);
        owner = payable(msg.sender);
        executor = _executor;
        if (msg.value > 0) {
            WETH.deposit{value: msg.value}();
        }
    }

    /**
        This function is called after your contract has received the flash loaned amount
     */
    function executeOperation(
        address[] calldata assets,
        uint256[] calldata amounts,
        uint256[] calldata premiums,
        address initiator,
        bytes calldata params
    ) external override returns (bool) {
        // Requires that the flashLoanGaurd be false  to continue
        require(flashLoanGuard == false);
        require(tx.origin == executor);
        flashLoanGuard = true;

        // This will destructure the encoded parameters data into necessary parts
        (address[] memory _targets, bytes[] memory _payloads) = abi.decode(
            params,
            (address[], bytes[])
        );

        for (uint256 i = 0; i < _targets.length; i++) {
            (bool _success, bytes memory _response) = _targets[i].call(
                _payloads[i]
            );
            require(_success);
            //_response;
        }

        // At the end of your logic above, this contract owes
        // the flashloaned amount + premiums.
        // Therefore ensure your contract has enough to repay
        // these amounts.

        for (uint256 i = 0; i < assets.length; i++) {
            uint256 amountOwing = amounts[i].add(premiums[i]);
            IERC20(assets[i]).approve(address(LENDING_POOL), amountOwing);
        }
        //        uint256 amountOwed = amount + premium;
        //        IERC20(asset).approve(address(POOL), amountOwed);

        return true;
    }

    function swapWithFlashLoan(
        address[] calldata _tokens,
        uint256 _ethAmountToCoinbase,
        uint256[] calldata _flashLoanAmounts,
        uint256[] calldata _modes,
        address[] memory _targets,
        bytes[] memory _payloads
    ) external payable onlyExecutor {
        //require(_targets.length == _payloads.length);

        bytes memory params = abi.encode(_targets, _payloads);

        uint256 _wethBalanceBefore = WETH.balanceOf(address(this));

        LENDING_POOL.flashLoan(
            address(this),
            _tokens,
            _flashLoanAmounts,
            _modes,
            address(this),
            params,
            0
        );

        uint256 _wethBalanceAfter = WETH.balanceOf(address(this));
        require(_wethBalanceAfter > _wethBalanceBefore + _ethAmountToCoinbase);
        if (_ethAmountToCoinbase == 0) return;

        uint256 _ethBalance = address(this).balance;
        if (_ethBalance < _ethAmountToCoinbase) {
            WETH.withdraw(_ethAmountToCoinbase - _ethBalance);
        }
        block.coinbase.transfer(_ethAmountToCoinbase);
        flashLoanGuard = false;
    }

    //Same as old logic
    // _wethAmountToFirstMarket is no longer used
    function swapWithoutFlashLoan(
        uint256 _ethAmountToCoinbase,
        address[] memory _targets,
        bytes[] memory _payloads
    ) external payable onlyExecutor {
        //require(_targets.length == _payloads.length);
        //Prevents flashloan from being called
        flashLoanGuard = true;
        uint256 _wethBalanceBefore = WETH.balanceOf(address(this));

        for (uint256 i = 0; i < _targets.length; i++) {
            (bool _success, bytes memory _response) = _targets[i].call(
                _payloads[i]
            );
            require(_success);
            //_response;
        }

        uint256 _wethBalanceAfter = WETH.balanceOf(address(this));
        require(_wethBalanceAfter > _wethBalanceBefore + _ethAmountToCoinbase);

        if (_ethAmountToCoinbase == 0) return;

        uint256 _ethBalance = address(this).balance;
        if (_ethBalance < _ethAmountToCoinbase) {
            WETH.withdraw(_ethAmountToCoinbase - _ethBalance);
        }
        block.coinbase.transfer(_ethAmountToCoinbase);
        flashLoanGuard = false;
    }

    function call(
        address payable _to,
        uint256 _value,
        bytes calldata _data
    ) external payable onlyOwner returns (bytes memory) {
        require(_to != address(0));
        (bool _success, bytes memory _result) = _to.call{value: _value}(_data);
        require(_success);
        return _result;
    }

    modifier onlyExecutor() {
        require(msg.sender == executor);
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    receive() external payable {}
}
