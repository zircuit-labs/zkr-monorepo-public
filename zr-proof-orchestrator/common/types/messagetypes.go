package types

// L2Block is the data that is produced to the L2BlocksQueue
type L2Block struct {
	Hash       Hash        `json:"hash"`
	ParentHash Hash        `json:"parent_hash"`
	Number     uint64      `json:"number"`
	Nonce      uint64      `json:"nonce"`
	Status     BlockStatus `json:"status"`
	BatchHash  *Hash       `json:"batch_hash,omitempty"`
}
