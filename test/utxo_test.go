package test

import (
	"log"
	"testing"

	"github.com/luozui/mini-bitcoin/lib/blockchain"
	"github.com/luozui/mini-bitcoin/lib/leveldb"
	"github.com/luozui/mini-bitcoin/lib/utxo"
)

func TestUTXO(t *testing.T) {
	var chain blockchain.Chain
	chain.Read()
	var utxoSet utxo.UTXO
	utxoSet = make(utxo.UTXO, 100000)
	log.Println(len(utxoSet))
	for _, block := range chain {
		if utxoSet.Addblock(&block) != nil {
			log.Println("add block error")
			t.Fail()
		}
		//block.Print(false)
		//log.Printf("hash:%v, height: %v\n", block.Hash, index+1)
	}
	if leveldb.DumpUTXO("utxo.ldb", &utxoSet) != nil {
		t.Fail()
	}
}
