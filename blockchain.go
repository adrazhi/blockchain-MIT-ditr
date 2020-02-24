package main

// Now that we built a simple block structure, let's
// take a look at the implementation of a simple blockchain
// In a blockchain blocks are stored sequentially and each
// insertion is linked to the previous one. Easy :)
type Blockchain struct {
	blocks []*Block
}

// Now that we created a simple blockchain we need to add blocks
// otherwise it would be useless
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.RootHash)
	bc.blocks = append(bc.blocks, newBlock)
}

// How can we create a new Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{InitialBlock()}}
}
