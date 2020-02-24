package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Blocks store valuable information
// For the purpose of this implemeentation we will include
// the following information fields
type Block struct {
	Timestamp     int64  // when was this block created
	Information   []byte // information that block contains
	NeighbourHash []byte // hash value of previous block
	RootHash      []byte // hash value of the root block
}

// Hashes are calculated using SHA256 (as an example)
// See line 25 below
func (b *Block) SetHash() {
	time := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join(
		[][]byte{
			b.NeighbourHash,
			b.Information,
			time},
		[]byte{})
	hash := sha256.Sum256(headers)
	b.RootHash = hash[:]
}

// Since we have said that the purpose of the hash
// is to be cryptographically hard so that new blocks cannot
// be easily added we still need a way to create the blocks
func NewBlock(data string, NeighbourHash []byte) *Block {
	created_block := &Block{time.Now().Unix(), []byte(data), NeighbourHash, []byte{}}
	created_block.SetHash()
	return created_block
}

// NewGenesisBlock creates and returns genesis Block
func InitialBlock() *Block {
	return NewBlock("Initial Block", []byte{})
}
