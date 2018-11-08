package blockchain

import (
	"../block"
	"../producer"
)

func (bc *Blockchain) AddBlock(data []byte) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := producer.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{producer.NewGenesisBlock()}}
}
