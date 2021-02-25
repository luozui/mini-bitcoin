package blockchain

// Chain , index + 1 is height
type Chain []Block

// Block is Bitcoin Block struct
type Block struct {
	Height             uint32
	MagicNumber        uint32
	Hash               [32]byte
	Size               uint32
	Header             Header
	TransactionCounter uint64
	Transactions       []Transaction
}

// Header of Block
type Header struct {
	Version           uint32
	PreviousBlockHash [32]byte
	MerkleRoot        []byte
	Timestamp         uint32
	Bits              uint32
	Nonce             uint32
}

// Input of Transaction
type Input struct {
	PreviousTransactionHash     [32]byte
	PreviousTransactionOutIndex uint32
	ScriptLength                uint64
	Script                      []byte
	DecodedScript               string
	SequenceNo                  []byte
}

// Output of Transaction
type Output struct {
	Value        uint64
	ScriptLength uint64
	Script       []byte
}

// Transaction of Block
type Transaction struct {
	Hash          [32]byte
	Version       uint32
	InputCounter  uint64
	Inputs        []Input
	OutputCounter uint64
	Outputs       []Output
	LockTime      uint32
}

// GetHeight - get block height
func (b *Block) GetHeight() uint32 {
	return b.Height
}
