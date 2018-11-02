package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 Coin to kian")
	bc.AddBlock("Send 2 more Coin to kian")

	for _, block := range bc.blocks {
		fmt.Printf("Pre. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}