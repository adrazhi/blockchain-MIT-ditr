package main

import (
	"fmt"
)

// This is the main functions of our program
// that calls other parts of it such as blockchain
// and blocks in order to execute
func main() {
	my_blockchain := NewBlockchain()

	// Add block transactions
	my_blockchain.AddBlock("Send 1 BTC to your learning facilitator")
	my_blockchain.AddBlock("Send 1 BTC to Professor John")
	my_blockchain.AddBlock("Send 1 BTC to Professor Sanchez")
	my_blockchain.AddBlock("Send 1 BTC to Alice")
	my_blockchain.AddBlock("Send 1 BTC to Bob")

	for _, block := range my_blockchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.NeighbourHash)
		fmt.Printf("Information: %s\n", block.Information)
		fmt.Printf("RootHash: %x\n", block.RootHash)
		fmt.Println()
	}
}
