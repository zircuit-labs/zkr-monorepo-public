package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/zircuit-labs/zkr-go-common/xerrors/stacktrace"
)

// BlockStatus strongly types the supported block status in our system.
type BlockStatus string

const (
	BlockStatusUnsafe             BlockStatus = "unsafe"
	BlockStatusSafe               BlockStatus = "safe"
	BlockStatusFinalized          BlockStatus = "finalized"
	BlockStatusBatched            BlockStatus = "batched"            // block was added to a real batch
	BlockStatusSerialized         BlockStatus = "serialized"         // block serialization complete: ready to be proven
	BlockStatusComponentsComplete BlockStatus = "componentscomplete" // block has all component proofs
	BlockStatusProven             BlockStatus = "proven"             // block has a proof
)

// SupportedBlockStatus is the set of all supported block status.
var SupportedBlockStatus = map[BlockStatus]bool{
	BlockStatusUnsafe:             true,
	BlockStatusSafe:               true,
	BlockStatusFinalized:          true,
	BlockStatusBatched:            true,
	BlockStatusSerialized:         true,
	BlockStatusComponentsComplete: true,
	BlockStatusProven:             true,
}

// ParseBlockStatus parses a string into a BlockStatus.
func ParseBlockStatus(t string) (BlockStatus, error) {
	if SupportedBlockStatus[BlockStatus(t)] {
		return BlockStatus(t), nil
	}
	return "", stacktrace.Wrap(fmt.Errorf("invalid block status %s", t))
}

// SQL custom type requirements
var _ driver.Valuer = (*BlockStatus)(nil)

func (h BlockStatus) Value() (driver.Value, error) {
	return string(h), nil
}

var _ sql.Scanner = (*BlockStatus)(nil)

func (h *BlockStatus) Scan(src interface{}) (err error) {
	switch src := src.(type) {
	case BlockStatus:
		*h = src
	case string:
		pt, err := ParseBlockStatus(src)
		if err != nil {
			return err
		}
		*h = pt
	case []byte:
		pt, err := ParseBlockStatus(string(src))
		if err != nil {
			return err
		}
		*h = pt
	default:
		return fmt.Errorf("unsupported block status: %T", src)
	}
	return nil
}

func (h BlockStatus) String() string {
	return string(h)
}
