package main

import (
	"./blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock([]byte("Send 1 LiT to Yang"))
	bc.AddBlock([]byte("Send 2 LiT to Yang"))

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", string(block.Data))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
