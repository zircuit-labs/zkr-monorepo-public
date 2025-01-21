package derive

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/crypto"

	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/predeploys"
	"github.com/zircuit-labs/zkr-monorepo-public/op-node/rollup"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/eth"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/solabi"
)

const (
	L1InfoFuncBedrockSignature           = "setL1BlockValues(uint64,uint64,uint256,bytes32,uint64,bytes32,uint256,uint256)"
	L1InfoFuncEcotoneSignature           = "setL1BlockValuesEcotone()"
	L1InfoExclusionsFuncSignature        = "setL1BlockValues(uint64,uint64,uint256,bytes32,uint64,bytes32,uint256,uint256,bytes)"
	L1InfoExclusionsFuncEcotoneSignature = "setL1BlockValuesEcotoneExclusions()"
	L1InfoArguments                      = 8
	L1InfoBedrockLen                     = 4 + 32*L1InfoArguments
	L1InfoEcotoneLen                     = 4 + 32*5 // after Ecotone upgrade, args are packed into 5 32-byte slots
)

var (
	L1InfoFuncBedrockBytes4           = crypto.Keccak256([]byte(L1InfoFuncBedrockSignature))[:4]
	L1InfoFuncEcotoneBytes4           = crypto.Keccak256([]byte(L1InfoFuncEcotoneSignature))[:4]
	L1InfoExclusionsFuncBytes4        = crypto.Keccak256([]byte(L1InfoExclusionsFuncSignature))[:4]
	L1InfoExclusionsFuncEcotoneBytes4 = crypto.Keccak256([]byte(L1InfoExclusionsFuncEcotoneSignature))[:4]
	L1InfoDepositerAddress            = common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001")
	L1BlockAddress                    = predeploys.L1BlockAddr
	ErrInvalidFormat                  = errors.New("invalid ecotone l1 block info format")
)

const (
	RegolithSystemTxGas = 1_000_000
)

// L1BlockInfo presents the information stored in a L1Block.setL1BlockValues call
type L1BlockInfo struct {
	*types.L1Info
}

func NewL1BlockInfo(l1Info *types.L1Info) *L1BlockInfo {
	return &L1BlockInfo{
		L1Info: l1Info,
	}
}

func L1InfoFromSystemTx(rollupCfg *rollup.Config, blockTime uint64, systemTx *types.Transaction) (*L1BlockInfo, error) {
	if systemTx.Type() != types.DepositTxType {
		return nil, fmt.Errorf("first payload tx has unexpected tx type: %d", systemTx.Type())
	}
	info, err := L1BlockInfoFromBytes(rollupCfg, blockTime, systemTx.Data())
	if err != nil {
		return nil, fmt.Errorf("failed to parse L1 info deposit tx from L2 block: %w", err)
	}
	return info, nil
}

func L1InfoFromBlock(rollupCfg *rollup.Config, block *types.Block) (*L1BlockInfo, error) {
	if len(block.Transactions()) == 0 {
		return nil, fmt.Errorf("l2 block is missing L1 info deposit tx, block hash: %s", block.Hash())
	}

	return L1InfoFromSystemTx(rollupCfg, block.Time(), block.Transactions()[0])
}

func L1InfoFromExecutionPayload(rollupCfg *rollup.Config, payload *eth.ExecutionPayload) (*L1BlockInfo, error) {
	if len(payload.Transactions) == 0 {
		return nil, fmt.Errorf("l2 block is missing L1 info deposit tx, block hash: %s", payload.BlockHash)
	}

	var tx types.Transaction
	if err := tx.UnmarshalBinary(payload.Transactions[0]); err != nil {
		return nil, fmt.Errorf("failed to decode first tx to read l1 info from: %w", err)
	}

	return L1InfoFromSystemTx(rollupCfg, uint64(payload.Timestamp), &tx)
}

// Bedrock Binary Format
// +---------+--------------------------+
// | Bytes   | Field                    |
// +---------+--------------------------+
// | 4       | Function signature       |
// | 32      | Number                   |
// | 32      | Time                     |
// | 32      | BaseFee                  |
// | 32      | BlockHash                |
// | 32      | SequenceNumber           |
// | 32      | BatcherHash              |
// | 32      | L1FeeOverhead            |
// | 32      | L1FeeScalar              |
// | 32      | DepositExclusions (offset)|
// | 32      | DepositExclusions (length)|
// | paddedN | DepositExclusions (data) |
// +---------+--------------------------+

func (info *L1BlockInfo) marshalBinaryBedrock() ([]byte, error) {
	exclusionBytes := new(bytes.Buffer)

	var solFuncSignature []byte
	totalLen := L1InfoBedrockLen
	// only use the new function if the deposit exclusions are not empty
	// this ensure backwards compatibility for derivation and makes the transaction
	// slightly smaller in the happy path (no exclusions)
	if info.DepositExclusions != nil && info.DepositExclusions.Count() > 0 {
		solFuncSignature = L1InfoExclusionsFuncBytes4
		if err := solabi.WriteBytes(exclusionBytes, info.DepositExclusions.MustBytes(), 0x120); err != nil {
			return nil, err
		}
		totalLen += exclusionBytes.Len()
	} else {
		solFuncSignature = L1InfoFuncBedrockBytes4
	}

	w := bytes.NewBuffer(make([]byte, 0, totalLen))
	if err := solabi.WriteSignature(w, solFuncSignature); err != nil {
		return nil, err
	}
	if err := solabi.WriteUint64(w, info.Number); err != nil {
		return nil, err
	}
	if err := solabi.WriteUint64(w, info.Time); err != nil {
		return nil, err
	}
	if err := solabi.WriteUint256(w, info.BaseFee); err != nil {
		return nil, err
	}
	if err := solabi.WriteHash(w, info.BlockHash); err != nil {
		return nil, err
	}
	if err := solabi.WriteUint64(w, info.SequenceNumber); err != nil {
		return nil, err
	}
	if err := solabi.WriteAddress(w, info.BatcherAddr); err != nil {
		return nil, err
	}
	if err := solabi.WriteEthBytes32(w, info.L1FeeOverhead); err != nil {
		return nil, err
	}
	if err := solabi.WriteEthBytes32(w, info.L1FeeScalar); err != nil {
		return nil, err
	}
	if _, err := w.Write(exclusionBytes.Bytes()); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (info *L1BlockInfo) unmarshalBinaryBedrock(data []byte) error {
	var solFuncSignature []byte
	var hasExclusions bool
	if len(data) == L1InfoBedrockLen {
		hasExclusions = false
		solFuncSignature = L1InfoFuncBedrockBytes4
	} else if len(data) > L1InfoBedrockLen {
		hasExclusions = true
		solFuncSignature = L1InfoExclusionsFuncBytes4
	} else {
		return fmt.Errorf("data is unexpected length: %d", len(data))
	}
	reader := bytes.NewReader(data)

	var err error
	if _, err := solabi.ReadAndValidateSignature(reader, solFuncSignature); err != nil {
		return err
	}
	if info.Number, err = solabi.ReadUint64(reader); err != nil {
		return err
	}
	if info.Time, err = solabi.ReadUint64(reader); err != nil {
		return err
	}
	if info.BaseFee, err = solabi.ReadUint256(reader); err != nil {
		return err
	}
	if info.BlockHash, err = solabi.ReadHash(reader); err != nil {
		return err
	}
	if info.SequenceNumber, err = solabi.ReadUint64(reader); err != nil {
		return err
	}
	if info.BatcherAddr, err = solabi.ReadAddress(reader); err != nil {
		return err
	}
	if info.L1FeeOverhead, err = solabi.ReadEthBytes32(reader); err != nil {
		return err
	}
	if info.L1FeeScalar, err = solabi.ReadEthBytes32(reader); err != nil {
		return err
	}

	if hasExclusions {
		var depositExclusionData []byte
		if depositExclusionData, err = solabi.ReadBytes(reader); err != nil {
			return err
		}
		info.DepositExclusions = MustBitmap(depositExclusionData)
	} else {
		// we are dealing with the old function signature without exclusions
		info.DepositExclusions = nil
	}

	if !solabi.EmptyReader(reader) {
		return errors.New("too many bytes")
	}
	return nil
}

// Ecotone Binary Format
// +---------+--------------------------+
// | Bytes   | Field                    |
// +---------+--------------------------+
// | 4       | Function signature       |
// | 4       | BaseFeeScalar            |
// | 4       | BlobBaseFeeScalar        |
// | 8       | SequenceNumber           |
// | 8       | Timestamp                |
// | 8       | L1BlockNumber            |
// | 32      | BaseFee                  |
// | 32      | BlobBaseFee              |
// | 32      | BlockHash                |
// | 32      | BatcherHash              |
// | 32      | ExclusionsLen            |
// | N       | ExclusionsData           |
// +---------+--------------------------+

func (info *L1BlockInfo) marshalBinaryEcotone() ([]byte, error) {
	exclusionBytes := new(bytes.Buffer)

	var solFuncSignature []byte
	totalLen := L1InfoEcotoneLen
	// only use the new function if the deposit exclusions are not empty
	// this ensure backwards compatibility for derivation and makes the transaction
	// slightly smaller in the happy path (no exclusions)
	if info.DepositExclusions != nil && info.DepositExclusions.Count() > 0 {
		solFuncSignature = L1InfoExclusionsFuncEcotoneBytes4
		if err := solabi.WriteBytesShort(exclusionBytes, info.DepositExclusions.MustBytes()); err != nil {
			return nil, err
		}
		totalLen += exclusionBytes.Len()
	} else {
		solFuncSignature = L1InfoFuncEcotoneBytes4
	}

	w := bytes.NewBuffer(make([]byte, 0, L1InfoEcotoneLen))
	if err := solabi.WriteSignature(w, solFuncSignature); err != nil {
		return nil, err
	}
	if err := binary.Write(w, binary.BigEndian, info.BaseFeeScalar); err != nil {
		return nil, err
	}
	if err := binary.Write(w, binary.BigEndian, info.BlobBaseFeeScalar); err != nil {
		return nil, err
	}
	if err := binary.Write(w, binary.BigEndian, info.SequenceNumber); err != nil {
		return nil, err
	}
	if err := binary.Write(w, binary.BigEndian, info.Time); err != nil {
		return nil, err
	}
	if err := binary.Write(w, binary.BigEndian, info.Number); err != nil {
		return nil, err
	}
	if err := solabi.WriteUint256(w, info.BaseFee); err != nil {
		return nil, err
	}
	blobBasefee := info.BlobBaseFee
	if blobBasefee == nil {
		blobBasefee = big.NewInt(1) // set to 1, to match the min blob basefee as defined in EIP-4844
	}
	if err := solabi.WriteUint256(w, blobBasefee); err != nil {
		return nil, err
	}
	if err := solabi.WriteHash(w, info.BlockHash); err != nil {
		return nil, err
	}
	// ABI encoding will perform the left-padding with zeroes to 32 bytes, matching the "batcherHash" SystemConfig format and version 0 byte.
	if err := solabi.WriteAddress(w, info.BatcherAddr); err != nil {
		return nil, err
	}
	if _, err := w.Write(exclusionBytes.Bytes()); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func (info *L1BlockInfo) unmarshalBinaryEcotone(data []byte) error {
	var solFuncSignature []byte
	var hasExclusions bool
	if len(data) == L1InfoEcotoneLen {
		hasExclusions = false
		solFuncSignature = L1InfoFuncEcotoneBytes4
	} else if len(data) > L1InfoEcotoneLen {
		hasExclusions = true
		solFuncSignature = L1InfoExclusionsFuncEcotoneBytes4
	} else {
		return fmt.Errorf("data is unexpected length: %d", len(data))
	}
	r := bytes.NewReader(data)

	var err error
	if _, err := solabi.ReadAndValidateSignature(r, solFuncSignature); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &info.BaseFeeScalar); err != nil {
		return ErrInvalidFormat
	}
	if err := binary.Read(r, binary.BigEndian, &info.BlobBaseFeeScalar); err != nil {
		return ErrInvalidFormat
	}
	if err := binary.Read(r, binary.BigEndian, &info.SequenceNumber); err != nil {
		return ErrInvalidFormat
	}
	if err := binary.Read(r, binary.BigEndian, &info.Time); err != nil {
		return ErrInvalidFormat
	}
	if err := binary.Read(r, binary.BigEndian, &info.Number); err != nil {
		return ErrInvalidFormat
	}
	if info.BaseFee, err = solabi.ReadUint256(r); err != nil {
		return err
	}
	if info.BlobBaseFee, err = solabi.ReadUint256(r); err != nil {
		return err
	}
	if info.BlockHash, err = solabi.ReadHash(r); err != nil {
		return err
	}
	// The "batcherHash" will be correctly parsed as address, since the version 0 and left-padding matches the ABI encoding format.
	if info.BatcherAddr, err = solabi.ReadAddress(r); err != nil {
		return err
	}

	if hasExclusions {
		var depositExclusionData []byte
		if depositExclusionData, err = solabi.ReadBytesShort(r); err != nil {
			return err
		}
		info.DepositExclusions = MustBitmap(depositExclusionData)
	} else {
		// we are dealing with the old function signature without exclusions
		info.DepositExclusions = nil
	}

	if !solabi.EmptyReader(r) {
		return errors.New("too many bytes")
	}
	return nil
}

// isEcotoneButNotFirstBlock returns whether the specified block is subject to the Ecotone upgrade,
// but is not the actiation block itself.
func isEcotoneButNotFirstBlock(rollupCfg *rollup.Config, l2BlockTime uint64) bool {
	return rollupCfg.IsEcotone(l2BlockTime) && !rollupCfg.IsEcotoneActivationBlock(l2BlockTime)
}

// L1BlockInfoFromBytes is the inverse of L1InfoDeposit, to see where the L2 chain is derived from
func L1BlockInfoFromBytes(rollupCfg *rollup.Config, l2BlockTime uint64, data []byte) (*L1BlockInfo, error) {
	return GetL1BlockInfo(isEcotoneButNotFirstBlock(rollupCfg, l2BlockTime), data)
}

func GetL1BlockInfo(Ecotone bool, data []byte) (*L1BlockInfo, error) {
	info := *NewL1BlockInfo(&types.L1Info{})
	if Ecotone {
		return &info, info.unmarshalBinaryEcotone(data)
	}
	return &info, info.unmarshalBinaryBedrock(data)
}

func L1BlockInfoFromParts(rollupCfg *rollup.Config, sysCfg eth.SystemConfig, seqNumber uint64, block eth.BlockInfo, l2BlockTime uint64, exclusions *types.Bitmap) (*L1BlockInfo, error) {
	l1BlockInfo := NewL1BlockInfo(&types.L1Info{
		Number:            block.NumberU64(),
		Time:              block.Time(),
		BaseFee:           block.BaseFee(),
		BlockHash:         block.Hash(),
		SequenceNumber:    seqNumber,
		BatcherAddr:       sysCfg.BatcherAddr,
		DepositExclusions: exclusions,
	})

	if isEcotoneButNotFirstBlock(rollupCfg, l2BlockTime) {
		l1BlockInfo.BlobBaseFee = block.BlobBaseFee()
		if l1BlockInfo.BlobBaseFee == nil {
			// The L2 spec states to use the MIN_BLOB_GASPRICE from EIP-4844 if not yet active on L1.
			l1BlockInfo.BlobBaseFee = big.NewInt(1)
		}
		scalars, err := sysCfg.EcotoneScalars()
		if err != nil {
			return nil, err
		}
		l1BlockInfo.BlobBaseFeeScalar = scalars.BlobBaseFeeScalar
		l1BlockInfo.BaseFeeScalar = scalars.BaseFeeScalar
	} else {
		l1BlockInfo.L1FeeOverhead = sysCfg.Overhead
		l1BlockInfo.L1FeeScalar = sysCfg.Scalar
	}

	return l1BlockInfo, nil
}

// L1InfoDeposit creates a L1 Info deposit transaction based on the L1 block,
// and the L2 block-height difference with the start of the epoch.
func L1InfoDeposit(rollupCfg *rollup.Config, sysCfg eth.SystemConfig, seqNumber uint64, block eth.BlockInfo, l2BlockTime uint64, exclusions *types.Bitmap) (*types.DepositTx, error) {
	l1BlockInfo, err := L1BlockInfoFromParts(rollupCfg, sysCfg, seqNumber, block, l2BlockTime, exclusions)
	if err != nil {
		return nil, fmt.Errorf("failed to create L1 block info: %w", err)
	}

	var data []byte
	if isEcotoneButNotFirstBlock(rollupCfg, l2BlockTime) {
		out, err := l1BlockInfo.marshalBinaryEcotone()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Ecotone l1 block info: %w", err)
		}
		data = out
	} else {
		out, err := l1BlockInfo.marshalBinaryBedrock()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Bedrock l1 block info: %w", err)
		}
		data = out
	}

	source := L1InfoDepositSource{
		L1BlockHash: block.Hash(),
		SeqNumber:   seqNumber,
	}
	// Set a very large gas limit with `IsSystemTransaction` to ensure
	// that the L1 Attributes Transaction does not run out of gas.
	out := &types.DepositTx{
		SourceHash:          source.SourceHash(),
		From:                L1InfoDepositerAddress,
		To:                  &L1BlockAddress,
		Mint:                nil,
		Value:               big.NewInt(0),
		Gas:                 150_000_000,
		IsSystemTransaction: true,
		Data:                data,
	}
	// With the regolith fork we disable the IsSystemTx functionality, and allocate real gas
	if rollupCfg.IsRegolith(l2BlockTime) {
		out.IsSystemTransaction = false
		out.Gas = RegolithSystemTxGas
	}
	return out, nil
}

// L1InfoDepositBytes returns a serialized L1-info attributes transaction.
func L1InfoDepositBytes(rollupCfg *rollup.Config, sysCfg eth.SystemConfig, seqNumber uint64, l1Info eth.BlockInfo, exclusions *types.Bitmap, l2BlockTime uint64) ([]byte, error) {
	dep, err := L1InfoDeposit(rollupCfg, sysCfg, seqNumber, l1Info, l2BlockTime, exclusions)
	if err != nil {
		return nil, fmt.Errorf("failed to create L1 info tx: %w", err)
	}
	l1Tx := types.NewTx(dep)
	opaqueL1Tx, err := l1Tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to encode L1 info tx: %w", err)
	}
	return opaqueL1Tx, nil
}
