package main

import (
	"blockchain"
	"fmt"
)


func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlocks() {
		fmt.Println("Data: %s\n", block.Data)
		fmt.Println("Hash: %s\n", block.Hash)
		fmt.Println("Prev Hash: %s\n", block.PrevHash)
	}
}
