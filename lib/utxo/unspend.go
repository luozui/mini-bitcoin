package utxo

import (
	"errors"

	"github.com/luozui/mini-bitcoin/lib/blockchain"
)

type UTXO map[UtxoKey]*UtxoVal

type UtxoKey struct {
	Type  uint32
	Txid  [32]byte
	Index uint32
}

type UtxoVal struct {
	Height     uint32
	Script     string
	ScriptLen  uint64
	Amount     uint64
	Iscoinbase bool
}

func (f *UTXO) Loadutxo() error {
	return nil
}

func (f *UTXO) Addblock(block *blockchain.Block) error {
	if (*f).check(block) == false {
		return errors.New("error")
	}
	for _, tx := range block.Transactions {
		if tx.InputCounter > 0 {
			(*f).del(tx.Inputs)
		}
		(*f).add(tx.Outputs, tx.Hash, block.GetHeight())
	}
	return nil
}

func (f *UTXO) del(txins []blockchain.Input) {
	for _, txin := range txins {
		delete((*f), UtxoKey{Txid: txin.PreviousTransactionHash, Index: txin.PreviousTransactionOutIndex})
	}
	return
}

func (f *UTXO) add(txouts []blockchain.Output, txid [32]byte, height uint32) {
	for i, txout := range txouts {
		(*f)[UtxoKey{Txid: txid, Index: uint32(i)}] = &UtxoVal{
			Height:     height,
			Script:     string(txout.Script),
			ScriptLen:  txout.ScriptLength,
			Amount:     txout.Value,
			Iscoinbase: i == 0,
		}
	}
}

// todo
func (f *UTXO) check(block *blockchain.Block) bool {
	return true
}
