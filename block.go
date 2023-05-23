package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Timestamp --> current timestamp when the block is created
// Data --> information in block
// PrevBlockHash --> hash of previous block
// Hash --> hash of the block

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Computing hash requires large amount of resources. Here we will just take block fields
// concatenate them and calculate the SHA-256 hash on the combination concatenated.

// Hashing the contents of block
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Creation of new blocks
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// in any blockchain, there must be at least one block,
// and such block, the first in the chain, is called genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
