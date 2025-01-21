// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { Deploy } from "scripts/Deploy.s.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Safe } from "safe-contracts/Safe.sol";

/// @notice Transfer the ownership of the proxy admin to a new address
contract TransferProxyAdminOwnership is Deploy {

    /// @inheritdoc Deploy
    function name() public pure virtual override returns (string memory name_) {
        name_ = "003-TransferProxyAdminOwnership";
    }

    function run() override pure public {
        revert("Not implemented. Use runSimulateFromMultisig.");
    }

    function runSimulateFromMultisig() public {
        address newOwner = vm.envAddress("NEW_PROXYADMIN_OWNER");
        transferOwnership(newOwner);
    }

    function transferOwnership(address _newOwner) broadcast internal {
        console.log("Transferring ownership to", _newOwner);
        ProxyAdmin proxyAdmin = ProxyAdmin(mustGetAddress("ProxyAdmin"));
        uint256 codeSize;
        assembly {
            codeSize := extcodesize(_newOwner)
        }
        require(codeSize > 0, "New owner does not seem to be a contract");


        Safe newSafe = Safe(payable(_newOwner));
        try newSafe.getThreshold() returns (uint256 threshold) {
            if (threshold <= 1) {
                console.log("Warning: New safe owner threshold is", threshold);
            }
        } catch {
            revert("New owner does not seem to be a safe contract");
        }

        proxyAdmin.transferOwnership(_newOwner);
    }
}

