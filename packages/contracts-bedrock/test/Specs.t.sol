// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { CommonTest } from "test/setup/CommonTest.sol";
import { Executables } from "scripts/Executables.sol";
import { console2 as console } from "forge-std/console2.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { SystemConfig } from "src/L1/SystemConfig.sol";
import { ForgeArtifacts, Abi, AbiEntry } from "scripts/ForgeArtifacts.sol";

/// @title Specification_Test
/// @dev Specifies common security properties of entrypoints to L1 contracts, including authorization and
///      pausability.
///      When adding new functions to the L1 system, the `setUp` function must be updated to document the security
///      properties of the new function. The `Spec` struct reppresents this documentation. However, this contract does
///      not actually test to verify these properties, only that a spec is defined.
contract Specification_Test is CommonTest {
    enum Role {
        NOAUTH,
        PROPOSER,
        CHALLENGER,
        SYSTEMCONFIGOWNER,
        MESSENGER,
        L1PROXYADMINOWNER,
        SUPERCHAINCONFIGMONITOR,
        SUPERCHAINCONFIGOPERATOR,
        SUPERCHAINCONFIGADMIN,
        L2OOSYSTEMOWNER
    }

    /// @notice Represents the specification of a function.
    /// @custom:field name     Contract name
    /// @custom:field sel      Function selector
    /// @custom:field auth     Specifies authentication as a requirement
    /// @custom:field pausable Specifies that the function is pausable
    struct Spec {
        string name;
        bytes4 sel;
        Role auth;
        bool pausable;
    }

    mapping(string => mapping(bytes4 => Spec)) specs;
    mapping(string => uint256) public numEntries;

    function setUp() public override {
        super.setUp();

        // L1CrossDomainMessenger
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("MESSAGE_VERSION()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("MIN_GAS_CALLDATA_OVERHEAD()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("OTHER_MESSENGER()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("PORTAL()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("RELAY_CALL_OVERHEAD()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("RELAY_CONSTANT_OVERHEAD()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("RELAY_GAS_CHECK_BUFFER()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("RELAY_RESERVED_GAS()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("baseGas(bytes,uint32)") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("failedMessages(bytes32)") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("initialize(address,address)") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("messageNonce()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("paused()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("otherMessenger()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("portal()") });
        _addSpec({
            _name: "L1CrossDomainMessenger",
            _sel: _getSel("relayMessage(uint256,address,address,uint256,uint256,bytes)"),
            _pausable: true
        });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("sendMessage(address,bytes,uint32)") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("successfulMessages(bytes32)") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("superchainConfig()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("version()") });
        _addSpec({ _name: "L1CrossDomainMessenger", _sel: _getSel("xDomainMessageSender()") });

        // L1ERC721Bridge
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("MESSENGER()") });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("OTHER_BRIDGE()") });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("bridgeERC721(address,address,uint256,uint32,bytes)") });
        _addSpec({
            _name: "L1ERC721Bridge",
            _sel: _getSel("bridgeERC721To(address,address,address,uint256,uint32,bytes)")
        });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("deposits(address,address,uint256)") });
        _addSpec({
            _name: "L1ERC721Bridge",
            _sel: _getSel("finalizeBridgeERC721(address,address,address,address,uint256,bytes)"),
            _pausable: true
        });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("messenger()") });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("otherBridge()") });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("version()") });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("superchainConfig()") });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("paused()") });
        _addSpec({ _name: "L1ERC721Bridge", _sel: _getSel("initialize(address,address)") });

        // L1StandardBridge
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("MESSENGER()") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("OTHER_BRIDGE()") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("accessController()") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("bridgeERC20(address,address,uint256,uint32,bytes)") });
        _addSpec({
            _name: "L1StandardBridge",
            _sel: _getSel("bridgeERC20To(address,address,address,uint256,uint32,bytes)")
        });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("bridgeETH(uint32,bytes)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("bridgeETHTo(address,uint32,bytes)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("deposits(address,address)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("erc20ThrottleDeposits(address)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("erc20ThrottleWithdrawals(address)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("ethThrottleDeposits()") });
        _addSpec({
            _name: "L1StandardBridge",
            _sel: _getSel("finalizeBridgeERC20(address,address,address,address,uint256,bytes)"),
            _auth: Role.MESSENGER,
            _pausable: true
        });
        _addSpec({
            _name: "L1StandardBridge",
            _sel: _getSel("finalizeBridgeETH(address,address,uint256,bytes)"),
            _auth: Role.MESSENGER,
            _pausable: true
        });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("getERC20ThrottleDepositsCredits(address,address)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("getERC20ThrottleWithdrawalsCredits(address)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("getEthThrottleDepositsCredits(address)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("initialize(address,address)") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("messenger()") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("otherBridge()") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("paused()") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("setErc20ThrottleDepositsMaxAmount(address,uint208,uint256)"), _auth: Role.SUPERCHAINCONFIGMONITOR });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("setErc20ThrottleDepositsPeriodLength(address,uint48)"), _auth: Role.SUPERCHAINCONFIGADMIN });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("setErc20ThrottleWithdrawalsMaxAmount(address,uint208,uint256)"), _auth: Role.SUPERCHAINCONFIGMONITOR });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("setErc20ThrottleWithdrawalsPeriodLength(address,uint48)"), _auth: Role.SUPERCHAINCONFIGADMIN });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("setEthThrottleDepositsMaxAmount(uint208,uint256)"), _auth: Role.SUPERCHAINCONFIGMONITOR });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("setEthThrottleDepositsPeriodLength(uint48)"), _auth: Role.SUPERCHAINCONFIGADMIN });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("superchainConfig()") });
        _addSpec({ _name: "L1StandardBridge", _sel: _getSel("version()") });

        // L2OutputOracle
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("CHALLENGER()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("FINALIZATION_PERIOD_SECONDS()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("L2_BLOCK_TIME()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("PROPOSER()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("SUBMISSION_INTERVAL()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("VERIFIER()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("challenger()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("computeL2Timestamp(uint256)") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("deleteL2Outputs(uint256)"), _auth: Role.CHALLENGER });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("finalizationPeriodSeconds()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("getL2Output(uint256)") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("getL2OutputAfter(uint256)") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("getL2OutputIndexAfter(uint256)") });
        _addSpec({
            _name: "L2OutputOracle",
            _sel: _getSel("initialize(uint256,uint256,uint256,uint256,address,address,address,uint256,address)")
        });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("l2BlockTime()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("latestBlockNumber()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("latestOutputIndex()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("nextBlockNumber()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("nextOutputIndex()") });
        _addSpec({
            _name: "L2OutputOracle",
            _sel: _getSel("proposeL2Output(bytes32,uint256,bytes32,uint256,bytes)"),
            _auth: Role.PROPOSER
        });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("proposer()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("setFinalizationPeriodSeconds(uint256)"), _auth: Role.L2OOSYSTEMOWNER });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("startingBlockNumber()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("startingTimestamp()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("submissionInterval()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("systemOwner()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("verifier()") });
        _addSpec({ _name: "L2OutputOracle", _sel: _getSel("version()") });



        // OptimismPortal
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("depositTransaction(address,uint256,uint64,bool,bytes)") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("donateETH()") });
        _addSpec({
            _name: "OptimismPortal",
            _sel: OptimismPortal.finalizeWithdrawalTransaction.selector,
            _pausable: true
        });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("finalizedWithdrawals(bytes32)") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("getEthThrottleDepositsCredits()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("getEthThrottleWithdrawalsCredits()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("guardian()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("initialize(address,address,address)") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("isOutputFinalized(uint256)") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("l2Oracle()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("l2Sender()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("minimumGasLimit(uint64)") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("params()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("paused()") });
        _addSpec({ _name: "OptimismPortal", _sel: OptimismPortal.proveWithdrawalTransaction.selector, _pausable: true });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("provenWithdrawals(bytes32)") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("superchainConfig()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("systemConfig()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("version()") });

        _addSpec({ _name: "OptimismPortal", _sel: _getSel("ethThrottleDeposits()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("ethThrottleWithdrawals()") });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("setEthThrottleDepositsMaxAmount(uint208,uint256)"), _auth: Role.SUPERCHAINCONFIGMONITOR });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("setEthThrottleDepositsPeriodLength(uint48)"), _auth: Role.SUPERCHAINCONFIGOPERATOR });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("setEthThrottleWithdrawalsMaxAmount(uint208,uint256)"), _auth: Role.SUPERCHAINCONFIGMONITOR });
        _addSpec({ _name: "OptimismPortal", _sel: _getSel("setEthThrottleWithdrawalsPeriodLength(uint48)"), _auth: Role.SUPERCHAINCONFIGOPERATOR });

        // ResourceMetering
        _addSpec({ _name: "ResourceMetering", _sel: _getSel("params()") });

        // SuperchainConfig
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("DEFAULT_ADMIN_ROLE()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("MONITOR_ROLE()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("OPERATOR_ROLE()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("PAUSED_SLOT()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("acceptDefaultAdminTransfer()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("beginDefaultAdminTransfer(address)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("cancelDefaultAdminTransfer()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("changeDefaultAdminDelay(uint48)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("defaultAdmin()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("defaultAdminDelay()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("defaultAdminDelayIncreaseWait()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("getRoleAdmin(bytes32)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("grantRole(bytes32,address)"), _auth: Role.SUPERCHAINCONFIGADMIN });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("guardian()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("hasMonitorCapabilities(address)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("hasOperatorCapabilities(address)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("hasRole(bytes32,address)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("initialize(address,bool)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("owner()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("pause(string)"), _auth: Role.SUPERCHAINCONFIGMONITOR });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("paused()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("pendingDefaultAdmin()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("pendingDefaultAdminDelay()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("renounceRole(bytes32,address)"), _auth: Role.SUPERCHAINCONFIGMONITOR });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("revokeRole(bytes32,address)"), _auth: Role.SUPERCHAINCONFIGADMIN });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("rollbackDefaultAdminDelay()") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("supportsInterface(bytes4)") });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("unpause()"), _auth: Role.SUPERCHAINCONFIGOPERATOR });
        _addSpec({ _name: "SuperchainConfig", _sel: _getSel("version()") });

        // SystemConfig
        _addSpec({ _name: "SystemConfig", _sel: _getSel("UNSAFE_BLOCK_SIGNER_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("START_BLOCK_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("VERSION()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("batcherHash()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("gasLimit()") });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.initialize.selector });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.minimumGasLimit.selector });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("overhead()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("owner()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("renounceOwnership()"), _auth: Role.SYSTEMCONFIGOWNER });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.resourceConfig.selector });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("scalar()") });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.setBatcherHash.selector, _auth: Role.SYSTEMCONFIGOWNER });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.setGasConfig.selector, _auth: Role.SYSTEMCONFIGOWNER });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.setGasLimit.selector, _auth: Role.SYSTEMCONFIGOWNER });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.setResourceConfig.selector, _auth: Role.SYSTEMCONFIGOWNER });
        _addSpec({
            _name: "SystemConfig",
            _sel: SystemConfig.setUnsafeBlockSigner.selector,
            _auth: Role.SYSTEMCONFIGOWNER
        });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("transferOwnership(address)"), _auth: Role.SYSTEMCONFIGOWNER });
        _addSpec({ _name: "SystemConfig", _sel: SystemConfig.unsafeBlockSigner.selector });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("version()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("l1CrossDomainMessenger()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("l1ERC721Bridge()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("l1StandardBridge()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("l2OutputOracle()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("optimismPortal()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("optimismMintableERC20Factory()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("batchInbox()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("startBlock()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("L1_CROSS_DOMAIN_MESSENGER_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("L1_ERC_721_BRIDGE_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("L1_STANDARD_BRIDGE_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("L2_OUTPUT_ORACLE_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("OPTIMISM_PORTAL_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("OPTIMISM_MINTABLE_ERC20_FACTORY_SLOT()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("BATCH_INBOX_SLOT()") });
        _addSpec({
            _name: "SystemConfig",
            _sel: _getSel("setGasConfigEcotone(uint32,uint32)"),
            _auth: Role.SYSTEMCONFIGOWNER
        });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("basefeeScalar()") });
        _addSpec({ _name: "SystemConfig", _sel: _getSel("blobbasefeeScalar()") });

        // ProxyAdmin
        _addSpec({
            _name: "ProxyAdmin",
            _sel: _getSel("changeProxyAdmin(address,address)"),
            _auth: Role.L1PROXYADMINOWNER
        });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("getProxyAdmin(address)") });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("getProxyImplementation(address)") });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("owner()") });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("proxyType(address)") });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("renounceOwnership()"), _auth: Role.L1PROXYADMINOWNER });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("setProxyType(address,uint8)"), _auth: Role.L1PROXYADMINOWNER });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("transferOwnership(address)"), _auth: Role.L1PROXYADMINOWNER });
        _addSpec({ _name: "ProxyAdmin", _sel: _getSel("upgrade(address,address)") });
        _addSpec({
            _name: "ProxyAdmin",
            _sel: _getSel("upgradeAndCall(address,address,bytes)"),
            _auth: Role.L1PROXYADMINOWNER
        });

        // Verifier
        _addSpec({ _name: "Verifier", _sel: _getSel("version()") });
    }

    /// @dev Computes the selector from a function signature.
    function _getSel(string memory _name) internal pure returns (bytes4) {
        return bytes4(keccak256(abi.encodePacked(_name)));
    }

    /// @dev Adds a spec for a function.
    function _addSpec(string memory _name, bytes4 _sel, Role _auth, bool _pausable) internal {
        specs[_name][_sel] = Spec({ name: _name, sel: _sel, auth: _auth, pausable: _pausable });
        numEntries[_name]++;
    }

    /// @dev Adds a spec for a function with no auth.
    function _addSpec(string memory _name, bytes4 _sel, bool _pausable) internal {
        _addSpec({ _name: _name, _sel: _sel, _auth: Role.NOAUTH, _pausable: _pausable });
    }

    /// @dev Adds a spec for a function with no pausability.
    function _addSpec(string memory _name, bytes4 _sel, Role _auth) internal {
        _addSpec({ _name: _name, _sel: _sel, _auth: _auth, _pausable: false });
    }

    /// @dev Adds a spec for a function with no auth and no pausability.
    function _addSpec(string memory _name, bytes4 _sel) internal {
        _addSpec({ _name: _name, _sel: _sel, _auth: Role.NOAUTH, _pausable: false });
    }

    /// @notice Ensures that there's an auth spec for every L1 contract function.
    function testContractAuth() public {
        string[] memory pathExcludes = new string[](1);
        pathExcludes[0] = "src/dispute/interfaces/*";
        Abi[] memory abis =
            ForgeArtifacts.getContractFunctionAbis("src/{L1,universal/ProxyAdmin.sol}", pathExcludes);

        for (uint256 i = 0; i < abis.length; i++) {
            string memory contractName = abis[i].contractName;
            assertEq(
                abis[i].entries.length,
                numEntries[contractName],
                string.concat("Specification_Test: invalid number of ABI entries for ", contractName)
            );

            for (uint256 j = 0; j < abis[i].entries.length; j++) {
                AbiEntry memory abiEntry = abis[i].entries[j];
                console.log(
                    "Checking auth spec for %s: %s(%x)", contractName, abiEntry.fnName, uint256(uint32(abiEntry.sel))
                );
                Spec memory spec = specs[contractName][abiEntry.sel];
                assertTrue(
                    spec.sel != bytes4(0),
                    string.concat(
                        "Specification_Test: missing spec definition for ", contractName, "::", abiEntry.fnName
                    )
                );
                assertEq(
                    abiEntry.sel,
                    spec.sel,
                    string.concat("Specification_Test: invalid ABI ", contractName, "::", abiEntry.fnName)
                );
            }
        }
    }
}
