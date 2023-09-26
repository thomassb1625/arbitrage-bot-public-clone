pragma solidity 0.6.12;

import {IERC20} from "protocol-v2/contracts/dependencies/openzeppelin/contracts/IERC20.sol";

interface Controller {
    function liquidate(address account, address product) external;
}

contract PerennialCheck {
    Controller controller = Controller(0x2d264EBDb6632A06A1726193D4d37FeF1E5dbDcd);

    function check(address[] memory accounts, address[] memory products) public returns (uint256[] memory) {
        uint256[] memory balances = new uint[](accounts.length);
        for (uint256 i = 0; i < accounts.length; i++) {
            (bool success, bytes memory resp ) = address(0x2d264EBDb6632A06A1726193D4d37FeF1E5dbDcd).call(
                abi.encodeWithSignature("liquidate(address,address)", accounts[i], products[i])
            );
            if (success) {
                uint256 balance = IERC20(0x605D26FBd5be761089281d5cec2Ce86eeA667109).balanceOf(address(this));
                balances[i] = balance;
                IERC20(0x605D26FBd5be761089281d5cec2Ce86eeA667109).transfer(address(0x9Dc07f920Ef2c70eF94FA45374A561CAC01f2617), balance);
            }
        }
        return balances;
    }
}
