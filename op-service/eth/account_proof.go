package eth

import (
	"bytes"
	"fmt"

	"github.com/holiman/uint256"
	zktrie "github.com/scroll-tech/zktrie/trie"
	zkt "github.com/scroll-tech/zktrie/types"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/common/hexutil"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/crypto"
	"github.com/zircuit-labs/l2-geth-public/ethdb/memorydb"
	"github.com/zircuit-labs/l2-geth-public/rlp"
	"github.com/zircuit-labs/l2-geth-public/trie"
)

type StorageProofEntry struct {
	Key   common.Hash     `json:"key"`
	Value hexutil.Big     `json:"value"`
	Proof []hexutil.Bytes `json:"proof"`
}

type AccountResult struct {
	AccountProof []hexutil.Bytes `json:"accountProof"`

	Address common.Address `json:"address"`
	Balance *hexutil.Big   `json:"balance"`
	// L1 code hash
	CodeHash common.Hash `json:"codeHash"`
	// L2 code hashes
	PoseidonCodeHash common.Hash    `json:"poseidonCodeHash"`
	KeccakCodeHash   common.Hash    `json:"keccakCodeHash"`
	CodeSize         hexutil.Uint64 `json:"codeSize"`
	Nonce            hexutil.Uint64 `json:"nonce"`
	StorageHash      common.Hash    `json:"storageHash"`

	// Optional
	StorageProof []StorageProofEntry `json:"storageProof,omitempty"`
}

// Verify an account (and optionally storage) proof from the getProof RPC. See https://eips.ethereum.org/EIPS/eip-1186
func (res *AccountResult) Verify(stateRoot common.Hash) error {
	// verify storage proof values, if any, against the storage trie root hash of the account
	for i, entry := range res.StorageProof {
		// load all MPT nodes into a DB
		db := memorydb.New()
		for j, encodedNode := range entry.Proof {
			nodeKey := encodedNode
			if len(encodedNode) >= 32 { // small MPT nodes are not hashed
				nodeKey = crypto.Keccak256(encodedNode)
			}
			if err := db.Put(nodeKey, encodedNode); err != nil {
				return fmt.Errorf("failed to load storage proof node %d of storage value %d into mem db: %w", j, i, err)
			}
		}
		path := crypto.Keccak256(entry.Key[:])
		val, err := trie.VerifyProof(res.StorageHash, path, db)
		if err != nil {
			verifyProofErr := err

			db = memorydb.New()
			if err := fillDBWithZKTrieNodes(db, entry.Proof); err != nil {
				return fmt.Errorf("failed to fill mem db with ZK trie nodes: %w", err)
			}

			val, err = trie.VerifyProofSMT(res.StorageHash, entry.Key[:], db)
			if err != nil {
				return fmt.Errorf("failed to verify storage value %d with key %s in storage trie %s: VerifyProof error: %v, VerifyProofSMT error: %w", i, entry.Key, res.StorageHash, verifyProofErr, err)
			}
			if val == nil && entry.Value.ToInt().Cmp(common.Big0) == 0 { // empty storage is zero by default
				continue
			}
			comparison := zkt.NewByte32FromBytes(entry.Value.ToInt().Bytes())
			if !bytes.Equal(val, comparison.Bytes()) {
				return fmt.Errorf("value %d in storage proof does not match proven value at key %s (path %x)", i, entry.Key, path)
			}
		} else {
			if val == nil && entry.Value.ToInt().Cmp(common.Big0) == 0 { // empty storage is zero by default
				continue
			}
			comparison, err := rlp.EncodeToBytes(entry.Value.ToInt().Bytes())
			if err != nil {
				return fmt.Errorf("failed to encode storage value %d with key %s (path %x) in storage trie %s: %w", i, entry.Key, path, res.StorageHash, err)
			}
			if !bytes.Equal(val, comparison) {
				return fmt.Errorf("value %d in storage proof does not match proven value at key %s (path %x)", i, entry.Key, path)
			}
		}
	}

	// L2 and L1 have different JSON tags.
	// Use codeHash as default
	codeHash := res.CodeHash
	if codeHash == (common.Hash{}) {
		codeHash = res.KeccakCodeHash
	}

	accountClaimed := []any{uint64(res.Nonce), res.Balance.ToInt().Bytes(), res.StorageHash, codeHash}
	accountClaimedValue, err := rlp.EncodeToBytes(accountClaimed)
	if err != nil {
		return fmt.Errorf("failed to encode account from retrieved values: %w", err)
	}

	// create a db with all account trie nodes
	db := memorydb.New()
	for i, encodedNode := range res.AccountProof {
		nodeKey := encodedNode
		if len(encodedNode) >= 32 { // small MPT nodes are not hashed
			nodeKey = crypto.Keccak256(encodedNode)
		}
		if err := db.Put(nodeKey, encodedNode); err != nil {
			return fmt.Errorf("failed to load account proof node %d into mem db: %w", i, err)
		}
	}
	path := crypto.Keccak256(res.Address[:])
	accountProofValue, err := trie.VerifyProof(stateRoot, path, db)
	if err != nil {
		verifyProofErr := err

		db = memorydb.New()
		if err := fillDBWithZKTrieNodes(db, res.AccountProof); err != nil {
			return fmt.Errorf("failed to fill mem db with ZK trie nodes: %w", err)
		}

		accountProofValue, err = trie.VerifyProofSMT(stateRoot, res.Address[:], db)
		if err != nil {
			return fmt.Errorf("failed to verify account value with key %s (path %x, SMT secure key: %x) in account trie %s: VerifyProof error: %v, VerifyProofSMT error: %w", res.Address, path, res.Address[:], stateRoot, verifyProofErr, err)
		}

		stateAccount := types.StateAccount{
			Nonce:            uint64(res.Nonce),
			Balance:          uint256.MustFromBig(res.Balance.ToInt()),
			Root:             common.HexToHash(res.StorageHash.Hex()),
			KeccakCodeHash:   res.KeccakCodeHash.Bytes(),
			PoseidonCodeHash: res.PoseidonCodeHash.Bytes(),
			CodeSize:         uint64(res.CodeSize),
		}

		stateAccMarshalled, _ := stateAccount.MarshalFields()
		accountClaimedValue = []byte("")
		for _, item := range stateAccMarshalled {
			accountClaimedValue = append(accountClaimedValue, item.Bytes()...)
		}

		if !bytes.Equal(accountClaimedValue, accountProofValue) {
			return fmt.Errorf("L1 RPC is tricking us, account proof does not match provided deserialized values:\n"+
				"  claimed:       %x\n"+
				"  proof:         %x\n", accountClaimedValue, accountProofValue)
		}
		return err
	}

	if !bytes.Equal(accountClaimedValue, accountProofValue) {
		return fmt.Errorf("L1 RPC is tricking us, account proof does not match provided deserialized values:\n"+
			"  claimed: %x\n"+
			"  proof:   %x", accountClaimedValue, accountProofValue)
	}
	return err
}

func fillDBWithZKTrieNodes(db *memorydb.Database, proof []hexutil.Bytes) error {
	for i, encodedNode := range proof {
		node, err := zktrie.DecodeSMTProof(encodedNode)
		if err != nil {
			return fmt.Errorf("failed to decode node %d from bytes: %w", i, err)
		}
		if node == nil {
			// skipping node with magic code
			continue
		}

		k, err := node.NodeHash()
		if err != nil {
			return fmt.Errorf("failed to hash proof node %d: %w", i, err)
		}

		if err := db.Put(k[:], encodedNode); err != nil {
			return fmt.Errorf("failed to load proof node %d into mem db: %w", i, err)
		}
	}
	return nil
}
