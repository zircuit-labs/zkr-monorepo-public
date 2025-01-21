// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { VmSafe } from "forge-std/Vm.sol";
import { Script } from "forge-std/Script.sol";

import { console2 as console } from "forge-std/console2.sol";
import { stdJson } from "forge-std/StdJson.sol";

import { Safe } from "safe-contracts/Safe.sol";
import { SafeProxyFactory } from "safe-contracts/proxies/SafeProxyFactory.sol";
import { Enum as SafeOps } from "safe-contracts/common/Enum.sol";

import { Deployer } from "scripts/Deployer.sol";

import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { L1StandardBridge } from "src/L1/L1StandardBridge.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { L1CrossDomainMessenger } from "src/L1/L1CrossDomainMessenger.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";
import { OptimismMintableERC20Factory } from "src/universal/OptimismMintableERC20Factory.sol";
import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";
import { SystemConfig } from "src/L1/SystemConfig.sol";
import { ResourceMetering } from "src/L1/ResourceMetering.sol";
import { Constants } from "src/libraries/Constants.sol";
import { L1ERC721Bridge } from "src/L1/L1ERC721Bridge.sol";
import { Verifier } from "src/L1/Verifier.sol";
import { MockVerifier } from "test/mocks/MockVerifier.sol";
import { StorageSetter } from "src/universal/StorageSetter.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Chains } from "scripts/Chains.sol";
import { Config } from "scripts/Config.sol";

import { ChainAssertions } from "scripts/ChainAssertions.sol";
import { Types } from "scripts/Types.sol";
import { LibStateDiff } from "scripts/libraries/LibStateDiff.sol";
import { EIP1967Helper } from "test/mocks/EIP1967Helper.sol";
import { ForgeArtifacts } from "scripts/ForgeArtifacts.sol";

/// @title Deploy
/// @notice Script used to deploy a bedrock system. The entire system is deployed within the `run` function.
///         To add a new contract to the system, add a public function that deploys that individual contract.
///         Then add a call to that function inside of `run`. Be sure to call the `save` function after each
///         deployment so that hardhat-deploy style artifacts can be generated using a call to `sync()`.
///         The `CONTRACT_ADDRESSES_PATH` environment variable can be set to a path that contains a JSON file full of
///         contract name to address pairs. That enables this script to be much more flexible in the way it is used.
///         This contract must not have constructor logic because it is set into state using `etch`.
contract Deploy is Deployer {
    using stdJson for string;

    ////////////////////////////////////////////////////////////////
    //                        Modifiers                           //
    ////////////////////////////////////////////////////////////////

    /// @notice Modifier that wraps a function in broadcasting.
    modifier broadcast() {
        vm.startBroadcast(msg.sender);
        _;
        vm.stopBroadcast();
    }

    /// @notice Modifier that will only allow a function to be called on devnet.
    modifier onlyDevnet() {
        uint256 chainid = block.chainid;
        if (chainid == Chains.LocalDevnet || chainid == Chains.GethDevnet) {
            _;
        }
    }

    /// @notice Modifier that will only allow a function to be called on a public
    ///         testnet or devnet.
    modifier onlyTestnetOrDevnet() {
        uint256 chainid = block.chainid;
        if (
            chainid == Chains.Sepolia || chainid == Chains.LocalDevnet
                || chainid == Chains.GethDevnet
        ) {
            _;
        }
    }

    /// @notice Modifier that wraps a function with statediff recording.
    ///         The returned AccountAccess[] array is then written to
    ///         the `snapshots/state-diff/<name>.json` output file.
    modifier stateDiff() {
        vm.startStateDiffRecording();
        _;
        VmSafe.AccountAccess[] memory accesses = vm.stopAndReturnStateDiff();
        console.log("Writing %d state diff account accesses to snapshots/state-diff/%s.json", accesses.length, name());
        string memory json = LibStateDiff.encodeAccountAccesses(accesses);
        string memory statediffPath = string.concat(vm.projectRoot(), "/snapshots/state-diff/", name(), ".json");
        vm.writeJson({ json: json, path: statediffPath });
    }

    ////////////////////////////////////////////////////////////////
    //                        Accessors                           //
    ////////////////////////////////////////////////////////////////

    /// @inheritdoc Deployer
    function name() public pure virtual override returns (string memory name_) {
        name_ = "Deploy";
    }

    /// @notice The create2 salt used for deployment of the contract implementations.
    ///         Using this helps to reduce config across networks as the implementation
    ///         addresses will be the same across networks when deployed with create2.
    function _implSalt() internal view returns (bytes32) {
        return keccak256(bytes(Config.implSalt()));
    }

    /// @notice Returns the proxy addresses. If a proxy is not found, it will have address(0).
    function _proxies() internal view returns (Types.ContractSet memory proxies_) {
        proxies_ = Types.ContractSet({
            L1CrossDomainMessenger: mustGetAddress("L1CrossDomainMessengerProxy"),
            L1StandardBridge: mustGetAddress("L1StandardBridgeProxy"),
            L2OutputOracle: mustGetAddress("L2OutputOracleProxy"),
            OptimismMintableERC20Factory: mustGetAddress("OptimismMintableERC20FactoryProxy"),
            OptimismPortal: mustGetAddress("OptimismPortalProxy"),
            SystemConfig: mustGetAddress("SystemConfigProxy"),
            L1ERC721Bridge: mustGetAddress("L1ERC721BridgeProxy"),
            SuperchainConfig: mustGetAddress("SuperchainConfigProxy"),
            Verifier: mustGetAddress("VerifierProxy")
        });
    }

    /// @notice Returns the proxy addresses, not reverting if any are unset.
    function _proxiesUnstrict() internal view returns (Types.ContractSet memory proxies_) {
        proxies_ = Types.ContractSet({
            L1CrossDomainMessenger: getAddress("L1CrossDomainMessengerProxy"),
            L1StandardBridge: getAddress("L1StandardBridgeProxy"),
            L2OutputOracle: getAddress("L2OutputOracleProxy"),
            OptimismMintableERC20Factory: getAddress("OptimismMintableERC20FactoryProxy"),
            OptimismPortal: getAddress("OptimismPortalProxy"),
            SystemConfig: getAddress("SystemConfigProxy"),
            L1ERC721Bridge: getAddress("L1ERC721BridgeProxy"),
            SuperchainConfig: getAddress("SuperchainConfigProxy"),
            Verifier: getAddress("VerifierProxy")
        });
    }

    ////////////////////////////////////////////////////////////////
    //            State Changing Helper Functions                 //
    ////////////////////////////////////////////////////////////////

    /// @notice Gets the address of the SafeProxyFactory and Safe singleton for use in deploying a new GnosisSafe.
    function _getSafeFactory() internal returns (SafeProxyFactory safeProxyFactory_, Safe safeSingleton_) {
        // These are the standard create2 deployed contracts. First we'll check if they are deployed,
        // if not we'll deploy new ones, though not at these addresses.
        address safeProxyFactory = 0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2;
        address safeSingleton = 0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552;

        safeProxyFactory.code.length == 0
            ? safeProxyFactory_ = new SafeProxyFactory()
            : safeProxyFactory_ = SafeProxyFactory(safeProxyFactory);

        safeSingleton.code.length == 0 ? safeSingleton_ = new Safe() : safeSingleton_ = Safe(payable(safeSingleton));
    }

    /// @notice Make a call from the Safe contract to an arbitrary address with arbitrary data
    function _callViaSafe(address _target, bytes memory _data) internal {
        Safe safe = Safe(mustGetAddress("SystemOwnerSafe"));

        // This is the signature format used the caller is also the signer.
        bytes memory signature = abi.encodePacked(uint256(uint160(msg.sender)), bytes32(0), uint8(1));

        safe.execTransaction({
            to: _target,
            value: 0,
            data: _data,
            operation: SafeOps.Operation.Call,
            safeTxGas: 0,
            baseGas: 0,
            gasPrice: 0,
            gasToken: address(0),
            refundReceiver: payable(address(0)),
            signatures: signature
        });
    }

    /// @notice Call from the Safe contract to the Proxy Admin's upgrade and call method
    function _upgradeAndCallViaSafe(address _proxy, address _implementation, bytes memory _innerCallData) internal {
        address proxyAdmin = mustGetAddress("ProxyAdmin");

        bytes memory data =
            abi.encodeCall(ProxyAdmin.upgradeAndCall, (payable(_proxy), _implementation, _innerCallData));

        _callViaSafe({ _target: proxyAdmin, _data: data });
    }

    /// @notice Call from the Safe contract to the Proxy Admin's upgrade method
    function _upgradeViaSafe(address _proxy, address _implementation) internal {
        address proxyAdmin = mustGetAddress("ProxyAdmin");

        bytes memory data =
            abi.encodeCall(ProxyAdmin.upgrade, (payable(_proxy), _implementation));

        _callViaSafe({ _target: proxyAdmin, _data: data });
    }

    /// @notice Transfer ownership of the ProxyAdmin contract to the final system owner
    function transferProxyAdminOwnership() public broadcast {
        ProxyAdmin proxyAdmin = ProxyAdmin(mustGetAddress("ProxyAdmin"));
        address owner = proxyAdmin.owner();
        address safe = mustGetAddress("SystemOwnerSafe");
        if (owner != safe) {
            proxyAdmin.transferOwnership(safe);
            console.log("ProxyAdmin ownership transferred to Safe at: %s", safe);
        }
    }

    /// @notice Transfer ownership of a Proxy to the ProxyAdmin contract
    ///         This is expected to be used in conjusting with deployERC1967ProxyWithOwner after setup actions
    ///         have been performed on the proxy.
    /// @param _name The name of the proxy to transfer ownership of.
    function transferProxyToProxyAdmin(string memory _name) public broadcast {
        Proxy proxy = Proxy(mustGetAddress(_name));
        address proxyAdmin = mustGetAddress("ProxyAdmin");
        proxy.changeAdmin(proxyAdmin);
        console.log("Proxy %s ownership transferred to ProxyAdmin at: %s", _name, proxyAdmin);
    }

    ////////////////////////////////////////////////////////////////
    //                    SetUp and Run                           //
    ////////////////////////////////////////////////////////////////

    /// @notice Deploy all of the L1 contracts necessary for a full Superchain with a single Op Chain.
    function run() public virtual {
        console.log("Deploying a fresh OP Stack including SuperchainConfig");
        _run();
    }

    function runWithStateDump() public {
        _run();
        vm.dumpState(Config.stateDumpPath(""));
    }

    /// @notice Deploy all L1 contracts and write the state diff to a file.
    function runWithStateDiff() public virtual stateDiff {
        _run();
    }

    /// @notice Internal function containing the deploy logic.
    function _run() internal {
        console.log("start of L1 Deploy!");
        deploySafe();
        console.log("deployed Safe!");
        setupSuperchain();
        console.log("set up superchain!");
        setupOpChain();
        console.log("set up op chain!");
    }

    ////////////////////////////////////////////////////////////////
    //           High Level Deployment Functions                  //
    ////////////////////////////////////////////////////////////////

    /// @notice Deploy a full system with a new SuperchainConfig
    ///         The Superchain system has 2 singleton contracts which lie outside of an OP Chain:
    ///         1. The SuperchainConfig contract
    function setupSuperchain() public {
        console.log("Setting up Superchain");

        // Deploy a new ProxyAdmin
        // This proxy will be used on the SuperchainConfig contract, as well as the contracts
        // in the OP Chain system.
        deployProxyAdmin();
        transferProxyAdminOwnership();

        // Deploy the SuperchainConfigProxy
        deployERC1967Proxy("SuperchainConfigProxy");
        deploySuperchainConfig();
        initializeSuperchainConfig();
    }

    /// @notice Deploy a new OP Chain, with an existing SuperchainConfig provided
    function setupOpChain() public {
        console.log("Deploying OP Chain");

        // Ensure that the requisite contracts are deployed
        mustGetAddress("SuperchainConfigProxy");
        mustGetAddress("SystemOwnerSafe");
        mustGetAddress("ProxyAdmin");

        deployProxies();
        deployImplementations();
        initializeImplementations();
    }

    /// @notice Deploy all of the proxies
    function deployProxies() public {
        console.log("Deploying proxies");

        deployERC1967Proxy("OptimismPortalProxy");
        deployERC1967Proxy("L2OutputOracleProxy");
        deployERC1967Proxy("SystemConfigProxy");
        deployERC1967Proxy("L1StandardBridgeProxy");
        deployERC1967Proxy("L1CrossDomainMessengerProxy");
        deployERC1967Proxy("OptimismMintableERC20FactoryProxy");
        deployERC1967Proxy("L1ERC721BridgeProxy");
        deployERC1967Proxy("VerifierProxy");
    }

    /// @notice Deploy all of the implementations
    function deployImplementations() public {
        console.log("Deploying implementations");
        deployL1CrossDomainMessenger();
        deployOptimismMintableERC20Factory();
        deploySystemConfig();
        deployL1StandardBridge();
        deployL1ERC721Bridge();
        deployVerifier();
        deployOptimismPortal();
        deployL2OutputOracle();
    }

    /// @notice Initialize all of the implementations
    function initializeImplementations() public {
        console.log("Initializing implementations");
        initializeSystemConfig();
        initializeL1StandardBridge();
        initializeL1ERC721Bridge();
        initializeOptimismMintableERC20Factory();
        initializeL1CrossDomainMessenger();
        initializeL2OutputOracle();
        initializeOptimismPortal();
        initializeVerifier();
    }

    ////////////////////////////////////////////////////////////////
    //              Non-Proxied Deployment Functions              //
    ////////////////////////////////////////////////////////////////

    /// @notice Deploy the Safe
    function deploySafe() public broadcast returns (address addr_) {
        console.log("Deploying Safe");
        (SafeProxyFactory safeProxyFactory, Safe safeSingleton) = _getSafeFactory();

        address[] memory signers = new address[](1);
        signers[0] = msg.sender;

        bytes memory initData = abi.encodeWithSelector(
            Safe.setup.selector, signers, 1, address(0), hex"", address(0), address(0), 0, address(0)
        );
        address safe = address(safeProxyFactory.createProxyWithNonce(address(safeSingleton), initData, block.timestamp));

        save("SystemOwnerSafe", address(safe));
        console.log("New SystemOwnerSafe deployed at %s", address(safe));
        addr_ = safe;
    }

    /// @notice Deploy the ProxyAdmin
    function deployProxyAdmin() public broadcast returns (address addr_) {
        console.log("Deploying ProxyAdmin");
        ProxyAdmin admin = new ProxyAdmin({ _owner: msg.sender });
        require(admin.owner() == msg.sender);

        save("ProxyAdmin", address(admin));
        console.log("ProxyAdmin deployed at %s", address(admin));
        addr_ = address(admin);
    }

    /// @notice Deploy the StorageSetter contract, used for upgrades.
    function deployStorageSetter(uint16 offset) public broadcast returns (address) {
        console.log("Deploying StorageSetter");
        bytes32 salt = keccak256("StorageSetter");
        StorageSetter setter;
        address deployAddr = vm.computeCreate2Address(salt, keccak256(bytes.concat(type(StorageSetter).creationCode, abi.encode(offset))));
        if (deployAddr.code.length != 0) {
            console.log("Using already deployed storage setter at", deployAddr);
            setter = StorageSetter(deployAddr);
        } else {
            setter = new StorageSetter{ salt: salt }(offset);
            console.log("StorageSetter deployed at: %s", address(setter));
        }
        require(deployAddr == address(setter), "Create2 address computation does not match deploy address");
        string memory version = setter.version();
        console.log("StorageSetter version: %s", version);
        return address(setter);
    }

    ////////////////////////////////////////////////////////////////
    //                Proxy Deployment Functions                  //
    ////////////////////////////////////////////////////////////////

    /// @notice Deploys an ERC1967Proxy contract with the ProxyAdmin as the owner.
    /// @param _name The name of the proxy contract to be deployed.
    /// @return addr_ The address of the deployed proxy contract.
    function deployERC1967Proxy(string memory _name) public returns (address addr_) {
        addr_ = deployERC1967ProxyWithOwner(_name, mustGetAddress("ProxyAdmin"));
    }

    /// @notice Deploys an ERC1967Proxy contract with a specified owner.
    /// @param _name The name of the proxy contract to be deployed.
    /// @param _proxyOwner The address of the owner of the proxy contract.
    /// @return addr_ The address of the deployed proxy contract.
    function deployERC1967ProxyWithOwner(
        string memory _name,
        address _proxyOwner
    )
        public
        broadcast
        returns (address addr_)
    {
        console.log(string.concat("Deploying ERC1967 proxy for ", _name));
        Proxy proxy = new Proxy({ _admin: _proxyOwner });

        require(EIP1967Helper.getAdmin(address(proxy)) == _proxyOwner);

        save(_name, address(proxy));
        console.log("   at %s", address(proxy));
        addr_ = address(proxy);
    }

    ////////////////////////////////////////////////////////////////
    //             Implementation Deployment Functions            //
    ////////////////////////////////////////////////////////////////

    /// @notice Deploy the SuperchainConfig contract
    function deploySuperchainConfig() public broadcast {
        SuperchainConfig superchainConfig = new SuperchainConfig{ salt: _implSalt() }();

        require(superchainConfig.guardian() == address(0xdead));
        save("SuperchainConfig", address(superchainConfig));

        require(loadInitializedSlot("SuperchainConfig") == 1, "SuperchainConfig is not initialized");
        console.log("SuperchainConfig deployed at %s", address(superchainConfig));
    }

    /// @notice Deploy the L1CrossDomainMessenger
    function deployL1CrossDomainMessenger() public broadcast returns (address addr_) {
        console.log("Deploying L1CrossDomainMessenger implementation");
        L1CrossDomainMessenger messenger = new L1CrossDomainMessenger{ salt: _implSalt() }();

        save("L1CrossDomainMessenger", address(messenger));
        console.log("L1CrossDomainMessenger deployed at %s", address(messenger));

        // Override the `L1CrossDomainMessenger` contract to the deployed implementation. This is necessary
        // to check the `L1CrossDomainMessenger` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.L1CrossDomainMessenger = address(messenger);
        ChainAssertions.checkL1CrossDomainMessenger({ _contracts: contracts, _vm: vm, _isProxy: false });

        addr_ = address(messenger);
    }

    /// @notice Deploy the OptimismPortal
    function deployOptimismPortal() public broadcast returns (address addr_) {
        console.log("Deploying OptimismPortal implementation");

        OptimismPortal portal = new OptimismPortal{ salt: _implSalt() }();

        save("OptimismPortal", address(portal));
        console.log("OptimismPortal deployed at %s", address(portal));

        // Override the `OptimismPortal` contract to the deployed implementation. This is necessary
        // to check the `OptimismPortal` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.OptimismPortal = address(portal);
        ChainAssertions.checkOptimismPortal({ _contracts: contracts, _cfg: cfg, _isProxy: false });

        require(loadInitializedSlot("OptimismPortal") == 1, "OptimismPortal is not initialized");

        addr_ = address(portal);
    }

    /// @notice Deploy the L2OutputOracle
    function deployL2OutputOracle() public broadcast returns (address addr_) {
        console.log("Deploying L2OutputOracle implementation");
        L2OutputOracle oracle = new L2OutputOracle{ salt: _implSalt() }();

        save("L2OutputOracle", address(oracle));
        console.log("L2OutputOracle deployed at %s", address(oracle));

        // Override the `L2OutputOracle` contract to the deployed implementation. This is necessary
        // to check the `L2OutputOracle` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.L2OutputOracle = address(oracle);
        ChainAssertions.checkL2OutputOracle({
            _contracts: contracts,
            _cfg: cfg,
            _l2OutputOracleStartingTimestamp: 0,
            _isProxy: false
        });

        addr_ = address(oracle);
    }

    /// @notice Deploy the OptimismMintableERC20Factory
    function deployOptimismMintableERC20Factory() public broadcast returns (address addr_) {
        console.log("Deploying OptimismMintableERC20Factory implementation");
        OptimismMintableERC20Factory factory = new OptimismMintableERC20Factory{ salt: _implSalt() }();

        save("OptimismMintableERC20Factory", address(factory));
        console.log("OptimismMintableERC20Factory deployed at %s", address(factory));

        // Override the `OptimismMintableERC20Factory` contract to the deployed implementation. This is necessary
        // to check the `OptimismMintableERC20Factory` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.OptimismMintableERC20Factory = address(factory);
        ChainAssertions.checkOptimismMintableERC20Factory({ _contracts: contracts, _isProxy: false });

        addr_ = address(factory);
    }

    /// @notice Deploy the SystemConfig
    function deploySystemConfig() public broadcast returns (address addr_) {
        console.log("Deploying SystemConfig implementation");
        SystemConfig config = new SystemConfig{ salt: _implSalt() }();

        save("SystemConfig", address(config));
        console.log("SystemConfig deployed at %s", address(config));

        // Override the `SystemConfig` contract to the deployed implementation. This is necessary
        // to check the `SystemConfig` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.SystemConfig = address(config);
        ChainAssertions.checkSystemConfig({ _contracts: contracts, _cfg: cfg, _isProxy: false });

        addr_ = address(config);
    }

    /// @notice Deploy the L1StandardBridge
    function deployL1StandardBridge() public broadcast returns (address addr_) {
        console.log("Deploying L1StandardBridge implementation");

        L1StandardBridge bridge = new L1StandardBridge{ salt: _implSalt() }();

        save("L1StandardBridge", address(bridge));
        console.log("L1StandardBridge deployed at %s", address(bridge));

        // Override the `L1StandardBridge` contract to the deployed implementation. This is necessary
        // to check the `L1StandardBridge` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.L1StandardBridge = address(bridge);
        ChainAssertions.checkL1StandardBridge({ _contracts: contracts, _isProxy: false });

        addr_ = address(bridge);
    }

    /// @notice Deploy the L1ERC721Bridge
    function deployL1ERC721Bridge() public broadcast returns (address addr_) {
        console.log("Deploying L1ERC721Bridge implementation");
        L1ERC721Bridge bridge = new L1ERC721Bridge{ salt: _implSalt() }();

        save("L1ERC721Bridge", address(bridge));
        console.log("L1ERC721Bridge deployed at %s", address(bridge));

        // Override the `L1ERC721Bridge` contract to the deployed implementation. This is necessary
        // to check the `L1ERC721Bridge` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.L1ERC721Bridge = address(bridge);

        ChainAssertions.checkL1ERC721Bridge({ _contracts: contracts, _isProxy: false });

        addr_ = address(bridge);
    }

    /// @notice Deploy the Verifier
    function deployVerifier() public broadcast returns (address addr_) {
        address verifier;
        console.log("Deploying Verifier implementation");

        uint256 chainid = block.chainid;
        if (chainid == Chains.LocalDevnet
            || chainid == Chains.GethDevnet
            || keccak256(abi.encodePacked(deploymentContext)) == keccak256(abi.encodePacked("mainnet-sepolia"))
            || keccak256(abi.encodePacked(deploymentContext)) == keccak256(abi.encodePacked("alphanet-sepolia"))) {
            console.log("WARNING: Deploying MockVerifier");
            verifier = address(new MockVerifier{ salt: _implSalt() }());
        } else {
            verifier = address(new Verifier{ salt: _implSalt() }());
        }

        save("Verifier", verifier);
        console.log("Verifier deployed at %s", verifier);

        addr_ = verifier;
    }

    ////////////////////////////////////////////////////////////////
    //                    Initialize Functions                    //
    ////////////////////////////////////////////////////////////////

    /// @notice Initialize the SuperchainConfig
    function initializeSuperchainConfig() public broadcast {
        address payable superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");
        address payable superchainConfig = mustGetAddress("SuperchainConfig");
        _upgradeAndCallViaSafe({
            _proxy: superchainConfigProxy,
            _implementation: superchainConfig,
            _innerCallData: abi.encodeCall(SuperchainConfig.initialize, (cfg.superchainConfigGuardian(), false))
        });

        ChainAssertions.checkSuperchainConfig({ _contracts: _proxiesUnstrict(), _cfg: cfg, _isPaused: false });
    }

    /// @notice Initialize the SystemConfig
    function initializeSystemConfig() public broadcast {
        console.log("Upgrading and initializing SystemConfig proxy");
        address systemConfigProxy = mustGetAddress("SystemConfigProxy");
        address systemConfig = mustGetAddress("SystemConfig");

        bytes32 batcherHash = bytes32(uint256(uint160(cfg.batchSenderAddress())));
        ResourceMetering.ResourceConfig memory resourceConfig = Constants.DEFAULT_RESOURCE_CONFIG();
        uint256 maxTx = cfg.maxTxPerBlock() / 2;
        if (maxTx > type(uint16).max)
            revert("derived max deposits per block does not fit into uint16");
        resourceConfig.maxTransactionLimit = uint16(maxTx);

        _upgradeAndCallViaSafe({
            _proxy: payable(systemConfigProxy),
            _implementation: systemConfig,
            _innerCallData: abi.encodeCall(
                SystemConfig.initialize,
                (
                    cfg.finalSystemOwner(),
                    cfg.basefeeScalar(),
                    cfg.blobbasefeeScalar(),
                    batcherHash,
                    uint64(cfg.l2GenesisBlockGasLimit()),
                    cfg.p2pSequencerAddress(),
                    Constants.DEFAULT_RESOURCE_CONFIG(),
                    cfg.batchInboxAddress(),
                    SystemConfig.Addresses({
                        l1CrossDomainMessenger: mustGetAddress("L1CrossDomainMessengerProxy"),
                        l1ERC721Bridge: mustGetAddress("L1ERC721BridgeProxy"),
                        l1StandardBridge: mustGetAddress("L1StandardBridgeProxy"),
                        l2OutputOracle: mustGetAddress("L2OutputOracleProxy"),
                        optimismPortal: mustGetAddress("OptimismPortalProxy"),
                        optimismMintableERC20Factory: mustGetAddress("OptimismMintableERC20FactoryProxy")
                    })
                )
            )
        });

        SystemConfig config = SystemConfig(systemConfigProxy);
        string memory version = config.version();
        console.log("SystemConfig version: %s", version);

        ChainAssertions.checkSystemConfig({ _contracts: _proxies(), _cfg: cfg, _isProxy: true });
    }

    /// @notice Initialize the L1StandardBridge
    function initializeL1StandardBridge() public broadcast {
        console.log("Upgrading and initializing L1StandardBridge proxy");
        address l1StandardBridgeProxy = mustGetAddress("L1StandardBridgeProxy");
        address l1StandardBridge = mustGetAddress("L1StandardBridge");
        address l1CrossDomainMessengerProxy = mustGetAddress("L1CrossDomainMessengerProxy");
        address superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");

        _upgradeAndCallViaSafe({
            _proxy: payable(l1StandardBridgeProxy),
            _implementation: l1StandardBridge,
            _innerCallData: abi.encodeCall(
                L1StandardBridge.initialize,
                (L1CrossDomainMessenger(l1CrossDomainMessengerProxy), SuperchainConfig(superchainConfigProxy))
            )
        });

        string memory version = L1StandardBridge(payable(l1StandardBridgeProxy)).version();
        console.log("L1StandardBridge version: %s", version);

        ChainAssertions.checkL1StandardBridge({ _contracts: _proxies(), _isProxy: true });
    }

    /// @notice Initialize the L1ERC721Bridge
    function initializeL1ERC721Bridge() public broadcast {
        console.log("Upgrading and initializing L1ERC721Bridge proxy");
        address l1ERC721BridgeProxy = mustGetAddress("L1ERC721BridgeProxy");
        address l1ERC721Bridge = mustGetAddress("L1ERC721Bridge");
        address l1CrossDomainMessengerProxy = mustGetAddress("L1CrossDomainMessengerProxy");
        address superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");

        _upgradeAndCallViaSafe({
            _proxy: payable(l1ERC721BridgeProxy),
            _implementation: l1ERC721Bridge,
            _innerCallData: abi.encodeCall(
                L1ERC721Bridge.initialize,
                (L1CrossDomainMessenger(payable(l1CrossDomainMessengerProxy)), SuperchainConfig(superchainConfigProxy))
            )
        });

        L1ERC721Bridge bridge = L1ERC721Bridge(l1ERC721BridgeProxy);
        string memory version = bridge.version();
        console.log("L1ERC721Bridge version: %s", version);

        ChainAssertions.checkL1ERC721Bridge({ _contracts: _proxies(), _isProxy: true });
    }

    /// @notice Ininitialize the OptimismMintableERC20Factory
    function initializeOptimismMintableERC20Factory() public broadcast {
        console.log("Upgrading and initializing OptimismMintableERC20Factory proxy");
        address optimismMintableERC20FactoryProxy = mustGetAddress("OptimismMintableERC20FactoryProxy");
        address optimismMintableERC20Factory = mustGetAddress("OptimismMintableERC20Factory");
        address l1StandardBridgeProxy = mustGetAddress("L1StandardBridgeProxy");

        _upgradeAndCallViaSafe({
            _proxy: payable(optimismMintableERC20FactoryProxy),
            _implementation: optimismMintableERC20Factory,
            _innerCallData: abi.encodeCall(OptimismMintableERC20Factory.initialize, (l1StandardBridgeProxy))
        });

        OptimismMintableERC20Factory factory = OptimismMintableERC20Factory(optimismMintableERC20FactoryProxy);
        string memory version = factory.version();
        console.log("OptimismMintableERC20Factory version: %s", version);

        ChainAssertions.checkOptimismMintableERC20Factory({ _contracts: _proxies(), _isProxy: true });
    }

    /// @notice initializeL1CrossDomainMessenger
    function initializeL1CrossDomainMessenger() public broadcast {
        console.log("Upgrading and initializing L1CrossDomainMessenger proxy");
        address l1CrossDomainMessengerProxy = mustGetAddress("L1CrossDomainMessengerProxy");
        address l1CrossDomainMessenger = mustGetAddress("L1CrossDomainMessenger");
        address superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");
        address optimismPortalProxy = mustGetAddress("OptimismPortalProxy");

        _upgradeAndCallViaSafe({
            _proxy: payable(l1CrossDomainMessengerProxy),
            _implementation: l1CrossDomainMessenger,
            _innerCallData: abi.encodeCall(
                L1CrossDomainMessenger.initialize,
                (SuperchainConfig(superchainConfigProxy), OptimismPortal(payable(optimismPortalProxy)))
            )
        });

        L1CrossDomainMessenger messenger = L1CrossDomainMessenger(l1CrossDomainMessengerProxy);
        string memory version = messenger.version();
        console.log("L1CrossDomainMessenger version: %s", version);

        ChainAssertions.checkL1CrossDomainMessenger({ _contracts: _proxies(), _vm: vm, _isProxy: true });

        require(
            loadInitializedSlot("L1CrossDomainMessenger") == 1, "L1CrossDomainMessengerProxy is not initialized"
        );
    }

    /// @notice Initialize the L2OutputOracle
    function initializeL2OutputOracle() public broadcast {
        console.log("Upgrading and initializing L2OutputOracle proxy");
        address l2OutputOracleProxy = mustGetAddress("L2OutputOracleProxy");
        address l2OutputOracle = mustGetAddress("L2OutputOracle");
        address verifierProxy = mustGetAddress("VerifierProxy");

        _upgradeAndCallViaSafe({
            _proxy: payable(l2OutputOracleProxy),
            _implementation: l2OutputOracle,
            _innerCallData: abi.encodeCall(
                L2OutputOracle.initialize,
                (
                    cfg.l2OutputOracleSubmissionInterval(),
                    cfg.l2BlockTime(),
                    cfg.l2OutputOracleStartingBlockNumber(),
                    cfg.l2OutputOracleStartingTimestamp(),
                    cfg.l2OutputOracleProposer(),
                    cfg.l2OutputOracleChallenger(),
                    verifierProxy,
                    cfg.finalizationPeriodSeconds(),
                    cfg.finalSystemOwner()
                )
            )
        });

        L2OutputOracle oracle = L2OutputOracle(l2OutputOracleProxy);
        string memory version = oracle.version();
        console.log("L2OutputOracle version: %s", version);

        ChainAssertions.checkL2OutputOracle({
            _contracts: _proxies(),
            _cfg: cfg,
            _l2OutputOracleStartingTimestamp: cfg.l2OutputOracleStartingTimestamp(),
            _isProxy: true
        });
    }

    /// @notice Initialize the OptimismPortal
    function initializeOptimismPortal() public broadcast {
        console.log("Upgrading and initializing OptimismPortal proxy");
        address optimismPortalProxy = mustGetAddress("OptimismPortalProxy");
        address optimismPortal = mustGetAddress("OptimismPortal");
        address l2OutputOracleProxy = mustGetAddress("L2OutputOracleProxy");
        address systemConfigProxy = mustGetAddress("SystemConfigProxy");
        address superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");

        _upgradeAndCallViaSafe({
            _proxy: payable(optimismPortalProxy),
            _implementation: optimismPortal,
            _innerCallData: abi.encodeCall(
                OptimismPortal.initialize,
                (
                    L2OutputOracle(l2OutputOracleProxy),
                    SystemConfig(systemConfigProxy),
                    SuperchainConfig(superchainConfigProxy)
                )
            )
        });

        OptimismPortal portal = OptimismPortal(payable(optimismPortalProxy));
        string memory version = portal.version();
        console.log("OptimismPortal version: %s", version);

        ChainAssertions.checkOptimismPortal({ _contracts: _proxies(), _cfg: cfg, _isProxy: true });

        require(loadInitializedSlot("OptimismPortal") == 1, "OptimismPortalProxy is not initialized");
    }

    /// @notice Initialize the Verifier
    function initializeVerifier() public broadcast {
        console.log("Upgrading and initializing Verifier proxy");
        address verifierProxy = mustGetAddress("VerifierProxy");
        address verifierImpl = mustGetAddress("Verifier");

        _upgradeViaSafe({
            _proxy: payable(verifierProxy),
            _implementation: verifierImpl
        });

        Verifier verifier = Verifier(verifierProxy);
        string memory version = verifier.version();
        console.log("Verifier version: %s", version);
    }
}
