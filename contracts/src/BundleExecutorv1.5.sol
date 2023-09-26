//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;
pragma experimental ABIEncoderV2;

import {FlashLoanSimpleReceiverBase} from "@aave-v3-core/contracts/flashloan/base/FlashLoanSimpleReceiverBase.sol";
import {IPoolAddressesProvider} from "@aave-v3-core/contracts/interfaces/IPoolAddressesProvider.sol";
import {IERC20} from "@aave-v3-core/contracts/dependencies/openzeppelin/contracts/IERC20.sol";

interface IWETH is IERC20 {
    function deposit() external payable;
    function withdraw(uint256) external;
}

contract FlashLoan is FlashLoanSimpleReceiverBase {

    // constructor(address _addressProvider)
    //     FlashLoanSimpleReceiverBase(IPoolAddressesProvider(_addressProvider))
    // {
    //     owner = payable(msg.sender);
    // }
    address payable public immutable owner;
    address public immutable executor;
    IWETH private immutable WETH;
    bool private flashLoanGuard;
    SwapData private swapData;
    //Cheaper to pack this into a struct
    struct SwapData {
        address[] targets;
        bytes[] payloads;
    }


    constructor(address _executor, address _weth, address _provider) payable
      FlashLoanSimpleReceiverBase(IPoolAddressesProvider(_provider)) {
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
        address asset,
        uint256 amount,
        uint256 premium,
        address initiator,
        bytes calldata params
    ) external override returns (bool) {
        //
        // This contract now has the funds requested.
        // Your logic goes here.
        
        //Loading swapData struct to memroy once so it is only read from storage once here as 
        //opposed to twice below
        SwapData memory _swapData = swapData;

        //Significantly cheaper with gas to load from storage and write to memory
        //prior to operating on data
        address[] memory _targets = _swapData.targets;
        bytes[] memory _payloads = _swapData.payloads;
        //Continue with our logic from swapping without flashloan 

        for (uint256 i = 0; i < _targets.length; i++) {
            (bool _success, bytes memory _response) = _targets[i].call(
                _payloads[i]
            );
            require(_success, "CE");
            _response;
        }

        // At the end of your logic above, this contract owes
        // the flashloaned amount + premiums.
        // Therefore ensure your contract has enough to repay
        // these amounts.

        // Approve the Pool contract allowance to *pull* the owed amount
        uint256 amountOwed = amount + premium;
        IERC20(asset).approve(address(POOL), amountOwed);

        return true;
    }

    function swapWithFlashLoan(
        address _token, 
        uint256 _ethAmountToCoinbase,
        uint256 _flashLoanAmount,
        address[] memory _targets,
        bytes[] memory _payloads
        ) external payable onlyExecutor {

        address receiverAddress = address(this);
        address asset = _token;
        uint256 amount = _flashLoanAmount;
        bytes memory params = "";
        uint16 referralCode = 0;
        require(_targets.length == _payloads.length, "WL");
        uint256 _wethBalanceBefore = WETH.balanceOf(address(this));

        //Setting storage variables
        swapData.payloads = _payloads;
        swapData.targets = _targets;

        POOL.flashLoanSimple(
            receiverAddress,
            asset,
            amount,
            params, 
            referralCode
        );

        uint256 _wethBalanceAfter = WETH.balanceOf(address(this));
        require(
            _wethBalanceAfter > _wethBalanceBefore + _ethAmountToCoinbase,
            "B0"
        );
        if (_ethAmountToCoinbase == 0) return;

        uint256 _ethBalance = address(this).balance;
        if (_ethBalance < _ethAmountToCoinbase) {
            WETH.withdraw(_ethAmountToCoinbase - _ethBalance);
        }
        block.coinbase.transfer(_ethAmountToCoinbase);
        flashLoanGuard = false;

    }

    //Same as old logic
    function swapWithoutFlashLoan(
        uint256 _wethAmountToFirstMarket,
        uint256 _ethAmountToCoinbase,
        address[] memory _targets,
        bytes[] memory _payloads
    ) external payable onlyExecutor {
        require(_targets.length == _payloads.length, "WL");
        //Prevents flashloan from being called
        flashLoanGuard = true;
        uint256 _wethBalanceBefore = WETH.balanceOf(address(this));
        WETH.transfer(_targets[0], _wethAmountToFirstMarket);
        for (uint256 i = 0; i < _targets.length; i++) {
            (bool _success, bytes memory _response) = _targets[i].call(
                _payloads[i]
            );
            require(_success, "CE");
            _response;
        }

        uint256 _wethBalanceAfter = WETH.balanceOf(address(this));
        require(
            _wethBalanceAfter > _wethBalanceBefore + _ethAmountToCoinbase,
            "B0"
        );
        if (_ethAmountToCoinbase == 0) return;

        uint256 _ethBalance = address(this).balance;
        if (_ethBalance < _ethAmountToCoinbase) {
            WETH.withdraw(_ethAmountToCoinbase - _ethBalance);
        }
        block.coinbase.transfer(_ethAmountToCoinbase);
        flashLoanGuard = false;
    }

    function getBalance(address _tokenAddress) external view returns (uint256) {
        return IERC20(_tokenAddress).balanceOf(address(this));
    }

    function withdraw(address _tokenAddress) external onlyOwner {
        IERC20 token = IERC20(_tokenAddress);
        token.transfer(msg.sender, token.balanceOf(address(this)));
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
