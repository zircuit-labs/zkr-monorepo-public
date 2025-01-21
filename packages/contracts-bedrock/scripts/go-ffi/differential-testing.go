package main

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"strconv"

	l1common "github.com/ethereum/go-ethereum/common"
	"github.com/zircuit-labs/l2-geth-public/accounts/abi"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/common/hexutil"
	"github.com/zircuit-labs/l2-geth-public/core/rawdb"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/crypto"
	"github.com/zircuit-labs/l2-geth-public/trie"
	"github.com/zircuit-labs/l2-geth-public/trie/triedb/hashdb"
	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/predeploys"
	"github.com/zircuit-labs/zkr-monorepo-public/op-chain-ops/crossdomain"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
)

// ABI types
var (
	// Plain dynamic dynBytes type
	dynBytes, _ = abi.NewType("bytes", "", nil)
	bytesArgs   = abi.Arguments{
		{Type: dynBytes},
	}

	// Plain fixed bytes32 type
	fixedBytes, _  = abi.NewType("bytes32", "", nil)
	fixedBytesArgs = abi.Arguments{
		{Type: fixedBytes},
	}

	uint32Type, _ = abi.NewType("uint32", "", nil)

	// Decoded nonce tuple (nonce, version)
	decodedNonce, _ = abi.NewType("tuple", "DecodedNonce", []abi.ArgumentMarshaling{
		{Name: "nonce", Type: "uint256"},
		{Name: "version", Type: "uint256"},
	})
	decodedNonceArgs = abi.Arguments{
		{Name: "encodedNonce", Type: decodedNonce},
	}

	// Decoded ecotone scalars (uint32, uint32)
	decodedScalars = abi.Arguments{
		{Name: "basefeeScalar", Type: uint32Type},
		{Name: "blobbasefeeScalar", Type: uint32Type},
	}

	// WithdrawalHash slot tuple (bytes32, bytes32)
	withdrawalSlot, _ = abi.NewType("tuple", "SlotHash", []abi.ArgumentMarshaling{
		{Name: "withdrawalHash", Type: "bytes32"},
		{Name: "zeroPadding", Type: "bytes32"},
	})
	withdrawalSlotArgs = abi.Arguments{
		{Name: "slotHash", Type: withdrawalSlot},
	}

	// Prove withdrawal inputs tuple (bytes32, bytes32, bytes32, bytes32, bytes[])
	proveWithdrawalInputs, _ = abi.NewType("tuple", "ProveWithdrawalInputs", []abi.ArgumentMarshaling{
		{Name: "worldRoot", Type: "bytes32"},
		{Name: "stateRoot", Type: "bytes32"},
		{Name: "outputRoot", Type: "bytes32"},
		{Name: "withdrawalHash", Type: "bytes32"},
		{Name: "proof", Type: "bytes[]"},
	})
	proveWithdrawalInputsArgs = abi.Arguments{
		{Name: "inputs", Type: proveWithdrawalInputs},
	}
)

func DiffTestUtils() {
	args := os.Args[2:]
	variant := args[0]

	// This command requires arguments
	if len(args) == 0 {
		panic("Error: No arguments provided")
	}

	switch variant {
	case "decodeVersionedNonce":
		// Parse input arguments
		input, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)

		// Decode versioned nonce
		nonce, version := crossdomain.DecodeVersionedNonce(input)

		// ABI encode output
		packArgs := struct {
			Nonce   *big.Int
			Version *big.Int
		}{
			nonce,
			version,
		}
		packed, err := decodedNonceArgs.Pack(&packArgs)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "encodeCrossDomainMessage":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		// Encode cross domain message
		encoded, err := encodeCrossDomainMessage(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error encoding cross domain message")

		// Pack encoded cross domain message
		packed, err := bytesArgs.Pack(&encoded)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "hashCrossDomainMessage":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		// Encode cross domain message
		encoded, err := encodeCrossDomainMessage(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error encoding cross domain message")

		// Hash encoded cross domain message
		hash := crypto.Keccak256Hash(encoded)

		// Pack hash
		packed, err := fixedBytesArgs.Pack(&hash)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "hashDepositTransaction":
		// Parse input arguments
		l1BlockHash := l1common.HexToHash(args[1])
		logIndex, ok := new(big.Int).SetString(args[2], 10)
		checkOk(ok)
		from := common.HexToAddress(args[3])
		to := common.HexToAddress(args[4])
		mint, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		value, ok := new(big.Int).SetString(args[6], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[7], 10)
		checkOk(ok)
		data := common.FromHex(args[8])

		// Create deposit transaction
		depositTx := makeDepositTx(from, to, value, mint, gasLimit, false, data, l1BlockHash, logIndex)

		// RLP encode deposit transaction
		encoded, err := types.NewTx(&depositTx).MarshalBinary()
		checkErr(err, "Error encoding deposit transaction")

		// Hash encoded deposit transaction
		hash := crypto.Keccak256Hash(encoded)

		// Pack hash
		packed, err := fixedBytesArgs.Pack(&hash)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "encodeDepositTransaction":
		// Parse input arguments
		from := common.HexToAddress(args[1])
		to := common.HexToAddress(args[2])
		value, ok := new(big.Int).SetString(args[3], 10)
		checkOk(ok)
		mint, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		isCreate := args[6] == "true"
		data := common.FromHex(args[7])
		l1BlockHash := l1common.HexToHash(args[8])
		logIndex, ok := new(big.Int).SetString(args[9], 10)
		checkOk(ok)

		depositTx := makeDepositTx(from, to, value, mint, gasLimit, isCreate, data, l1BlockHash, logIndex)

		// RLP encode deposit transaction
		encoded, err := types.NewTx(&depositTx).MarshalBinary()
		checkErr(err, "Failed to RLP encode deposit transaction")
		// Pack rlp encoded deposit transaction
		packed, err := bytesArgs.Pack(&encoded)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "hashWithdrawal":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		// Hash withdrawal
		hash, err := hashWithdrawal(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error hashing withdrawal")

		// Pack hash
		packed, err := fixedBytesArgs.Pack(&hash)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "hashOutputRootProof":
		// Parse input arguments
		version := common.HexToHash(args[1])
		stateRoot := common.HexToHash(args[2])
		messagePasserStorageRoot := common.HexToHash(args[3])
		latestBlockHash := common.HexToHash(args[4])

		// Hash the output root proof
		hash, err := hashOutputRootProof(version, stateRoot, messagePasserStorageRoot, latestBlockHash)
		checkErr(err, "Error hashing output root proof")

		// Pack hash
		packed, err := fixedBytesArgs.Pack(&hash)
		checkErr(err, "Error encoding output")

		fmt.Print(hexutil.Encode(packed))
	case "getProveWithdrawalTransactionInputs":
		// Parse input arguments
		nonce, ok := new(big.Int).SetString(args[1], 10)
		checkOk(ok)
		sender := common.HexToAddress(args[2])
		target := common.HexToAddress(args[3])
		value, ok := new(big.Int).SetString(args[4], 10)
		checkOk(ok)
		gasLimit, ok := new(big.Int).SetString(args[5], 10)
		checkOk(ok)
		data := common.FromHex(args[6])

		wdHash, err := hashWithdrawal(nonce, sender, target, value, gasLimit, data)
		checkErr(err, "Error hashing withdrawal")

		// Compute the storage slot the withdrawalHash will be stored in
		slot := struct {
			WithdrawalHash common.Hash
			ZeroPadding    common.Hash
		}{
			WithdrawalHash: wdHash,
			ZeroPadding:    common.Hash{},
		}
		packed, err := withdrawalSlotArgs.Pack(&slot)
		checkErr(err, "Error packing withdrawal slot")

		// Compute the storage slot the withdrawalHash will be stored in
		hash := crypto.Keccak256Hash(packed)

		// Create a secure trie for state
		state, err := trie.NewStateTrie(
			trie.TrieID(types.EmptyLegacyTrieRootHash),
			trie.NewDatabase(rawdb.NewMemoryDatabase(), &trie.Config{HashDB: hashdb.Defaults}),
		)
		checkErr(err, "Error creating secure trie")

		// Put a "true" bool in the storage slot
		err = state.UpdateStorage(common.Address{}, hash.Bytes(), []byte{0x01})
		checkErr(err, "Error updating storage")

		// Create a secure trie for the world state
		world, err := trie.NewStateTrie(
			trie.TrieID(types.EmptyLegacyTrieRootHash),
			trie.NewDatabase(rawdb.NewMemoryDatabase(), &trie.Config{HashDB: hashdb.Defaults}),
		)
		checkErr(err, "Error creating secure trie")

		// Put the put the rlp encoded account in the world trie
		account := types.StateAccount{
			Nonce:   0,
			Balance: common.U2560,
			Root:    state.Hash(),
		}
		writer := new(bytes.Buffer)
		checkErr(account.EncodeRLP(writer), "Error encoding account")
		err = world.UpdateStorage(common.Address{}, predeploys.L2ToL1MessagePasserAddr.Bytes(), writer.Bytes())
		checkErr(err, "Error updating storage")

		// Get the proof
		var proof proofList
		checkErr(state.Prove(predeploys.L2ToL1MessagePasserAddr.Bytes(), &proof), "Error getting proof")

		// Get the output root
		outputRoot, err := hashOutputRootProof(common.Hash{}, world.Hash(), state.Hash(), common.Hash{})
		checkErr(err, "Error hashing output root proof")

		// Pack the output
		output := struct {
			WorldRoot      common.Hash
			StateRoot      common.Hash
			OutputRoot     common.Hash
			WithdrawalHash common.Hash
			Proof          proofList
		}{
			WorldRoot:      world.Hash(),
			StateRoot:      state.Hash(),
			OutputRoot:     outputRoot,
			WithdrawalHash: wdHash,
			Proof:          proof,
		}
		packed, err = proveWithdrawalInputsArgs.Pack(&output)
		checkErr(err, "Error encoding output")

		// Print the output
		fmt.Print(hexutil.Encode(packed[32:]))
	case "encodeScalarEcotone":
		basefeeScalar, err := strconv.ParseUint(args[1], 10, 32)
		checkErr(err, "Error decoding basefeeScalar")
		blobbasefeeScalar, err := strconv.ParseUint(args[2], 10, 32)
		checkErr(err, "Error decoding blobbasefeeScalar")

		encoded := eth.EncodeScalar(eth.EcotoneScalars{
			BaseFeeScalar:     uint32(basefeeScalar),
			BlobBaseFeeScalar: uint32(blobbasefeeScalar),
		})
		fmt.Print(hexutil.Encode(encoded[:]))
	case "decodeScalarEcotone":
		scalar := common.HexToHash(args[1])
		scalars, err := eth.DecodeScalar([32]byte(scalar[:]))
		checkErr(err, "Error decoding scalar")

		packed, err := decodedScalars.Pack(scalars.BaseFeeScalar, scalars.BlobBaseFeeScalar)
		checkErr(err, "Error encoding output")
		fmt.Print(hexutil.Encode(packed))
	default:
		panic(fmt.Errorf("Unknown command: %s", args[0]))
	}
}
