package producer

import (
  "../block"
  "../consensus/pow"
  "time"
)


func NewBlock(data []byte, prevBlockHash []byte) *block.Block {
	block := block.Block{time.Now().Unix(), data, prevBlockHash, []byte{}, 0}
	proofOfWork := pow.NewProofOfWork(&block)
	nonce, hash := proofOfWork.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

func NewGenesisBlock() *block.Block {
	return NewBlock([]byte("Genesis Block"), []byte{})
}
