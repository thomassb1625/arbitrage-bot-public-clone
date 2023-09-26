// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
// import "src/ExecutorV3.yul";
import "src/YulDeployer.sol";

interface UniswapOptimized {
    function owner() external returns (address);
}

interface IERC20 {
    function balanceOf(address) external view returns (uint256);
}

interface Pair {
    function price0CumulativeLast() external view returns (uint256);
    function price1CumulativeLast() external view returns (uint256);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

contract ExecutorTest is Test {
    YulDeployer yulDeployer = new YulDeployer();
    address owner = address(0xCAC9325A179B4dbB2FfbF76BD3524EEF966fd81A);
    address univ3Owner = address(0xC49995D326b9438CF4b0AFA1045f2bAD777Cf37F);
    address ownerWithdraw = address(0x98827093418C47C7fD5F161eC5B8DF1dA474bD9E);
    address adversary = address(0xE0A53921b3FB731c4EeB0D3Fb9f2Ad1057F6E650);
    UniswapOptimized arbRouter;
    IERC20 weth = IERC20(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);

    function setUp() public {
        arbRouter = UniswapOptimized(yulDeployer.deployContract("ExecutorV3"));
    }


    function testDodo() public{
        vm.deal(address(univ3Owner), 100000000 wei);

        vm.startPrank(address(univ3Owner), address(univ3Owner));
        address pair0 = address(0x4585FE77225b41b697C938B018E2Ac67Ac5a20c0); // WBTC_WETH
        uint72 ethIn  = 4.37773784e18;
        deal(address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2), address(arbRouter), ethIn);
        (bool status,) = address(arbRouter).call{value: 1024+256+128 +8+ 2}(
            abi.encodePacked(ethIn)
        );
        assertTrue(status, "expectSucess: call reverted");
    }


    function testV3ThreeSwap() public{
        vm.deal(address(univ3Owner), 100000000 wei);

        vm.startPrank(address(univ3Owner), address(univ3Owner));
        deal(address(0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599), address(arbRouter), 1);
        deal(address(0xdAC17F958D2ee523a2206206994597C13D831ec7), address(arbRouter), 1);

        address pair0 = address(0x4585FE77225b41b697C938B018E2Ac67Ac5a20c0); // WBTC_WETH
        uint72 ethIn  = 4.37773784e18;
        deal(address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2), address(arbRouter), ethIn);
        address pair1 = address(0x9Db9e0e53058C89e5B94e29621a205198648425B); //WBTC_USDT
        address pair1Token = address(0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599);
        uint112 pair1In =  0.30543617e8;

        address pair2 = address(0x11b815efB8f581194ae79006d24E0d814B7697F6); //WETH_USDT
        address pair2Token = address(0xdAC17F958D2ee523a2206206994597C13D831ec7);
        uint112 pair2In = 4864.05776e6;

        uint256 checkpointGasLeft = gasleft();
        (bool status,) = address(arbRouter).call{value: 1024+256+128 +8+ 2}(
            abi.encodePacked(ethIn,pair0, pair1, pair1Token, pair1In, pair2, pair2Token, pair2In)
        );
        uint256 checkpointGasLeft2 = gasleft();
        uint256 gasDelta = checkpointGasLeft - checkpointGasLeft2 - 100;
        emit log_named_uint("Gas used", gasDelta + 21000 + 1100);

        emit log_named_uint("balance After", weth.balanceOf(address(arbRouter)));
        assertTrue(status, "expectSucess: call reverted");
    }

    function testOtherVariableSwap() public {
        vm.deal(address(univ3Owner), 100000000 wei);

        vm.startPrank(address(univ3Owner), address(univ3Owner));

        address pair0 = address(0x1CFe0C85F16E03E91a4A27187511104150736d88); //uniV2 GOO_WETH 0.05%
        uint72 ethIn = 0.120712072213576111e18; //WETH IN
        uint112 pair0Out = 202853.858060518e9; //OUT
        deal(address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2), address(arbRouter), ethIn);
        emit log_named_uint("balance before", weth.balanceOf(address(arbRouter)));

        address pair1 = address(0xB3c7EFC915782F1F2a71EEeB009fFc2C87f8395d); // uniswapv3 GOO_WETH 0.3%
        address pair1Token = address(0xa01710cA98E4d66fD8D2044B3437C024e7A64D76);
        uint112 pair1In =  202853.858060518e9;

        //vm.expectRevert();
        uint256 checkpointGasLeft = gasleft();
        (bool status,) = address(arbRouter).call{value: 256+128 + 1}(
            abi.encodePacked(ethIn, pair0Out, pair0, pair1, pair1Token, pair1In)
        );
        uint256 checkpointGasLeft2 = gasleft();
        uint256 gasDelta = checkpointGasLeft - checkpointGasLeft2 - 100;
        emit log_named_uint("Gas used", gasDelta + 21000 + 1100);

        emit log_named_uint("balance After", weth.balanceOf(address(arbRouter)));
        assertTrue(status, "expectSucess: call reverted");
    }

    function testVariableSwaps() public {
        vm.deal(address(univ3Owner), 1000000000 wei);

        vm.startPrank(address(univ3Owner), address(univ3Owner));

        address pair0 = address(0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc); //uniV2 USDC_WETH
        uint72 pair0Amount = 1e18; //WETH IN
        deal(address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2), address(arbRouter), pair0Amount);
        emit log_named_uint("balance before", weth.balanceOf(address(arbRouter)));

        uint112 pair0Out = 1100e6; // USDC OUT
        address pair1 = address(0xAE461cA67B15dc8dc81CE7615e0320dA1A9aB8D5); // univ2 DAI_USDC
        uint112 pair1out = 1050e18;

        address pair2 = address(0xC2e9F25Be6257c210d7Adf0D4Cd6E3E881ba25f8); // univ3 DAI_WETH
        address pair2TokenIn = address(0x6B175474E89094C44Da98b954EedeAC495271d0F); //DAI
        uint112 pair2In = 999e18; // DAI IN

        //vm.expectRevert();
        // numswap = numbr of swaps - 1
        //Offset is wrong, you are not considering the 0 swap
        //        SWAP3 UNIV3   + SWAP 2 Zeroforone   + intial Weth transfer(yes) + numSwaps(3)
        uint256 value = 2048 + 1024 + 32 + 8 * 0 + 2;
        uint256 checkpointGasLeft = gasleft();
        (bool status,) = address(arbRouter).call{value: value}(
            abi.encodePacked(pair0Amount, pair0Out, pair0, pair1, pair1out, pair2, pair2TokenIn, pair2In)
        );

        uint256 checkpointGasLeft2 = gasleft();
        uint256 gasDelta = checkpointGasLeft - checkpointGasLeft2 - 100;
        emit log_named_uint("Gas used", gasDelta);

        assertTrue(status, "expectSucess: call  reverted");
    }

    function testWrongCaller() public {
        vm.deal(address(adversary), 1000 wei);

        vm.startPrank(adversary, adversary);

        address pair0 = address(0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640); //uniV3 WETH_USDC
        uint72 pair0Amount = 3e18; //WETH IN
        deal(address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2), address(arbRouter), pair0Amount);

        address pair1 = address(0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc); // uniswap WETH_USDC
        uint72 pair1Amount = 2e18; // WETH OUT
        vm.expectRevert();
        (bool status,) = address(arbRouter).call{value: 0}(abi.encodePacked(pair0, pair0Amount, pair1, pair1Amount));

        assertTrue(status, "expectRevert: call did not revert");
    }

    function testRevert() public {
        //Test if contract reverts if uniswap can't output at least ethOut
        vm.deal(address(owner), 1000 wei);

        vm.startPrank(address(owner), address(owner));

        address pair0 = address(0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc); // uniswap WETH_USDC
        uint72 pair0Amount = 3e18;
        deal(address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2), address(arbRouter), pair0Amount);

        address pair1 = address(0x397FF1542f962076d0BFE58eA045FfA2d347ACa0); // sushi WETH_USDC
        uint112 pair1Amount = 3000e6;
        uint72 ethOut = 3.00000000000001e18;

        vm.expectRevert();
        (bool status,) =
            address(arbRouter).call{value: 17 wei}(abi.encodePacked(pair0, pair0Amount, pair1, pair1Amount, ethOut));
        assertTrue(status, "expectRevert: call did not revert");
    }

    // forge test -vvvvv --fork-url "https://nd-505-088-180.p2pify.com/507c3d8ba5ea481bc66a09c4194f8de0";
    function testRoute1() public {
        vm.deal(address(owner), 1000 wei);

        vm.startPrank(address(owner));

        address pair0 = address(0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc); // uniswap WETH_USDC
        uint72 pair0Amount = 1e18;
        deal(address(weth), address(arbRouter), pair0Amount);

        address pair1 = address(0x397FF1542f962076d0BFE58eA045FfA2d347ACa0); // sushi WETH_USDC
        uint256 lastPrice = Pair(pair1).price0CumulativeLast();

        (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast) = Pair(pair1).getReserves();

        emit log_named_uint("price", pair0Amount / (reserve1 / reserve0));
        uint112 pair1Amount = pair0Amount / (reserve1 / reserve0) * 996 / 1000 - 1;
        uint72 ethOut = 1e18 * 95 / 100;

        (bool success, bytes memory ret) =
            address(arbRouter).call{value: 17 wei}(abi.encodePacked(pair0, pair0Amount, pair1, pair1Amount, ethOut));

        assertTrue(weth.balanceOf(address(arbRouter)) != pair0Amount);
        assertTrue(weth.balanceOf(address(arbRouter)) >= ethOut);

        //Should revert because balance after is smaller than before
        emit log_named_uint("eth out", weth.balanceOf(address(arbRouter)));
        emit log_bytes(ret);

        assertTrue(success);
    }

    function testWithdraw() public {
        vm.deal(address(ownerWithdraw), 1000 wei);
        vm.deal(address(arbRouter), 100000000 wei);

        assertTrue(IERC20(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2).balanceOf(address(ownerWithdraw)) == 0);
        vm.startPrank(address(ownerWithdraw));
        uint80 val = 1e18;
        deal(address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2), address(arbRouter), val);

        (bool success, bytes memory ret) = address(arbRouter).call{value: 0 wei}(abi.encodePacked(val));
        emit log_named_uint("call status", success == true ? 1 : 0);
        emit log_bytes(ret);

        assertTrue(IERC20(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2).balanceOf(address(arbRouter)) == 0);
        assertTrue(IERC20(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2).balanceOf(address(ownerWithdraw)) == val);
        emit log_named_uint("eth balance owner", address(ownerWithdraw).balance);
        emit log_named_uint("weth balance owner", IERC20(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2).balanceOf(address(ownerWithdraw)));
        assertTrue(success);
    }
}
