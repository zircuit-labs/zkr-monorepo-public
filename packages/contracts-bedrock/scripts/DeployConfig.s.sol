// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import { Script } from "forge-std/Script.sol";
import { console2 as console } from "forge-std/console2.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { Executables } from "scripts/Executables.sol";
import { Chains } from "scripts/Chains.sol";
import { VmSafe } from "forge-std/Vm.sol";

/// @title DeployConfig
/// @notice Represents the configuration required to deploy the system. It is expected
///         to read the file from JSON. A future improvement would be to have fallback
///         values if they are not defined in the JSON themselves.
contract DeployConfig is Script {
    string internal _json;

    address public finalSystemOwner;
    address public superchainConfigGuardian;
    uint256 public l1ChainID;
    uint256 public l2ChainID;
    uint256 public l2BlockTime;
    uint256 public maxSequencerDrift;
    uint256 public sequencerWindowSize;
    uint256 public channelTimeout;
    address public p2pSequencerAddress;
    address public batchInboxAddress;
    address public batchSenderAddress;
    uint256 public l2OutputOracleSubmissionInterval;
    int256 internal _l2OutputOracleStartingTimestamp;
    uint256 public l2OutputOracleStartingBlockNumber;
    address public l2OutputOracleProposer;
    address public l2OutputOracleChallenger;
    uint256 public finalizationPeriodSeconds;
    bool public fundDevAccounts;
    address public proxyAdminOwner;
    address public baseFeeVaultRecipient;
    uint256 public baseFeeVaultMinimumWithdrawalAmount;
    uint256 public baseFeeVaultWithdrawalNetwork;
    address public l1FeeVaultRecipient;
    uint256 public l1FeeVaultMinimumWithdrawalAmount;
    uint256 public l1FeeVaultWithdrawalNetwork;
    address public sequencerFeeVaultRecipient;
    uint256 public sequencerFeeVaultMinimumWithdrawalAmount;
    uint256 public sequencerFeeVaultWithdrawalNetwork;
    uint256 public l2GenesisBlockGasLimit;
    uint32 public basefeeScalar;
    uint32 public blobbasefeeScalar;
    uint256 public l2GenesisBlockBaseFeePerGas;
    uint256 public eip1559Denominator;
    uint256 public eip1559Elasticity;
    uint256 public systemConfigStartBlock;
    uint256 public maxTxPerBlock;

    function read(string memory _path) public {
        console.log("DeployConfig: reading file %s", _path);
        try vm.readFile(_path) returns (string memory data) {
            _json = data;
        } catch {
            require(false, string.concat("Cannot find deploy config file at ", _path));
        }

        finalSystemOwner = stdJson.readAddress(_json, "$.finalSystemOwner");
        superchainConfigGuardian = stdJson.readAddress(_json, "$.superchainConfigGuardian");
        l1ChainID = stdJson.readUint(_json, "$.l1ChainID");
        l2ChainID = stdJson.readUint(_json, "$.l2ChainID");
        l2BlockTime = stdJson.readUint(_json, "$.l2BlockTime");
        maxSequencerDrift = stdJson.readUint(_json, "$.maxSequencerDrift");
        sequencerWindowSize = stdJson.readUint(_json, "$.sequencerWindowSize");
        channelTimeout = stdJson.readUint(_json, "$.channelTimeout");
        p2pSequencerAddress = stdJson.readAddress(_json, "$.p2pSequencerAddress");
        batchInboxAddress = stdJson.readAddress(_json, "$.batchInboxAddress");
        batchSenderAddress = stdJson.readAddress(_json, "$.batchSenderAddress");
        l2OutputOracleSubmissionInterval = stdJson.readUint(_json, "$.l2OutputOracleSubmissionInterval");
        _l2OutputOracleStartingTimestamp = stdJson.readInt(_json, "$.l2OutputOracleStartingTimestamp");
        l2OutputOracleStartingBlockNumber = stdJson.readUint(_json, "$.l2OutputOracleStartingBlockNumber");
        l2OutputOracleProposer = stdJson.readAddress(_json, "$.l2OutputOracleProposer");
        l2OutputOracleChallenger = stdJson.readAddress(_json, "$.l2OutputOracleChallenger");
        finalizationPeriodSeconds = stdJson.readUint(_json, "$.finalizationPeriodSeconds");
        fundDevAccounts = _readOr(_json, "$.fundDevAccounts", false);
        proxyAdminOwner = stdJson.readAddress(_json, "$.proxyAdminOwner");
        baseFeeVaultRecipient = stdJson.readAddress(_json, "$.baseFeeVaultRecipient");
        baseFeeVaultMinimumWithdrawalAmount = stdJson.readUint(_json, "$.baseFeeVaultMinimumWithdrawalAmount");
        baseFeeVaultWithdrawalNetwork = stdJson.readUint(_json, "$.baseFeeVaultWithdrawalNetwork");
        l1FeeVaultRecipient = stdJson.readAddress(_json, "$.l1FeeVaultRecipient");
        l1FeeVaultMinimumWithdrawalAmount = stdJson.readUint(_json, "$.l1FeeVaultMinimumWithdrawalAmount");
        l1FeeVaultWithdrawalNetwork = stdJson.readUint(_json, "$.l1FeeVaultWithdrawalNetwork");
        sequencerFeeVaultRecipient = stdJson.readAddress(_json, "$.sequencerFeeVaultRecipient");
        sequencerFeeVaultMinimumWithdrawalAmount = stdJson.readUint(_json, "$.sequencerFeeVaultMinimumWithdrawalAmount");
        sequencerFeeVaultWithdrawalNetwork = stdJson.readUint(_json, "$.sequencerFeeVaultWithdrawalNetwork");
        l2GenesisBlockGasLimit = stdJson.readUint(_json, "$.l2GenesisBlockGasLimit");
        l2GenesisBlockBaseFeePerGas = stdJson.readUint(_json, "$.l2GenesisBlockBaseFeePerGas");
        basefeeScalar = uint32(_readOr(_json, "$.gasPriceOracleBaseFeeScalar", 1368));
        blobbasefeeScalar = uint32(_readOr(_json, "$.gasPriceOracleBlobBaseFeeScalar", 810949));
        eip1559Denominator = stdJson.readUint(_json, "$.eip1559Denominator");
        eip1559Elasticity = stdJson.readUint(_json, "$.eip1559Elasticity");
        systemConfigStartBlock = stdJson.readUint(_json, "$.systemConfigStartBlock");
        maxTxPerBlock = stdJson.readUint(_json, "$.maxTxPerBlock");
    }

    function l1StartingBlockTag() public returns (bytes32) {
        try vm.parseJsonBytes32(_json, "$.l1StartingBlockTag") returns (bytes32 tag) {
            return tag;
        } catch {
            try vm.parseJsonString(_json, "$.l1StartingBlockTag") returns (string memory tag) {
                return _getBlockByTag(tag);
            } catch {
                try vm.parseJsonUint(_json, "$.l1StartingBlockTag") returns (uint256 tag) {
                    return _getBlockByTag(vm.toString(tag));
                } catch { }
            }
        }
        revert("l1StartingBlockTag must be a bytes32, string or uint256 or cannot fetch l1StartingBlockTag");
    }

    function l2OutputOracleStartingTimestamp() public returns (uint256) {
        if (_l2OutputOracleStartingTimestamp < 0) {
            bytes32 tag = l1StartingBlockTag();
            string[] memory cmd = new string[](3);
            cmd[0] = Executables.bash;
            cmd[1] = "-c";
            cmd[2] = string.concat("cast block ", vm.toString(tag), " --json | ", Executables.jq, " .timestamp");
            bytes memory res = vm.ffi(cmd);
            return stdJson.readUint(string(res), "");
        }
        return uint256(_l2OutputOracleStartingTimestamp);
    }

    /// @notice Allow the `fundDevAccounts` config to be overridden.
    function setFundDevAccounts(bool _fundDevAccounts) public {
        fundDevAccounts = _fundDevAccounts;
    }

    function _getBlockByTag(string memory _tag) internal returns (bytes32) {
        string[] memory cmd = new string[](3);
        cmd[0] = Executables.bash;
        cmd[1] = "-c";
        cmd[2] = string.concat("cast block ", _tag, " --json | ", Executables.jq, " -r .hash");
        console.log("cmd %s %s %s", cmd[0], cmd[1], cmd[2]);
        bytes memory res = vm.ffi(cmd);
        VmSafe.FfiResult memory res2 = vm.tryFfi(cmd);
        console.log(vm.toString(res2.stderr));
        console.log(vm.toString(res2.stdout));
        console.log(vm.toString(res2.exitCode));
        console.log(vm.toString(res));
        return abi.decode(res, (bytes32));
    }

    function _readOr(string memory json, string memory key, bool defaultValue) internal view returns (bool) {
        return vm.keyExists(json, key) ? stdJson.readBool(json, key) : defaultValue;
    }

    function _readOr(string memory json, string memory key, uint256 defaultValue) internal view returns (uint256) {
        return vm.keyExists(json, key) ? stdJson.readUint(json, key) : defaultValue;
    }

    function _readOr(string memory json, string memory key, address defaultValue) internal view returns (address) {
        return vm.keyExists(json, key) ? stdJson.readAddress(json, key) : defaultValue;
    }

    function _readOr(
        string memory json,
        string memory key,
        string memory defaultValue
    )
        internal
        view
        returns (string memory)
    {
        return vm.keyExists(json, key) ? stdJson.readString(json, key) : defaultValue;
    }
}
