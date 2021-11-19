package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestmap     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

//block hash calculator
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestmap, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//new block
func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}}
	block.SetHash()
	return block
}

//instance block
func NewGenesisBlock() *Block {
	return NewBlock("Instance Block", []byte{})
}
