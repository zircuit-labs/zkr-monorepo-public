package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/zircuit-labs/l2-geth-public/common"
)

// Hash is a wrapper around common.Hash to allow for custom JSON marshalling.
type Hash struct {
	Hash common.Hash `json:"hash"`
}

// UnmarshalText parses a hash in hex syntax.
func (h *Hash) UnmarshalText(input []byte) error {
	return h.Hash.UnmarshalText(input)
}

// UnmarshalJSON parses a hash in hex syntax.
func (h *Hash) UnmarshalJSON(input []byte) error {
	return h.Hash.UnmarshalJSON(input)
}

// MarshalText returns the hex representation of h.
func (h *Hash) MarshalText() ([]byte, error) {
	return h.Hash.MarshalText()
}

func (h *Hash) IsZero() bool {
	return h.Hash == common.Hash{}
}

func (h Hash) String() string {
	return h.Hash.Hex()
}

// Short returns the last 6 characters of the hash.
func (h Hash) Short() string {
	hx := h.Hash.Hex()
	return hx[len(hx)-6:]
}

func StringToHash(h string) Hash {
	return Hash{Hash: common.HexToHash(h)}
}

// SQL custom type requirements
var _ driver.Valuer = (*Hash)(nil)

func (h Hash) Value() (driver.Value, error) {
	return h.String(), nil
}

var _ sql.Scanner = (*Hash)(nil)

func (h *Hash) Scan(src interface{}) (err error) {
	switch src := src.(type) {
	case Hash:
		*h = src
	case string:
		h.Hash = common.HexToHash(src)
	case []byte:
		h.Hash = common.HexToHash(string(src))
	case nil:
		h.Hash = common.Hash{}
	default:
		return fmt.Errorf("unsupported data type: %T", src)
	}
	return nil
}

type (
	HashSorter []Hash
)

func (n HashSorter) Len() int           { return len(n) }
func (n HashSorter) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n HashSorter) Less(i, j int) bool { return n[i].String() < n[j].String() }
