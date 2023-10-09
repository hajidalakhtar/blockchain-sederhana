package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

// mmebuat struck block
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// membuat fungsi untuk menghitung hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// membuat fungsi untuk membuat block baru dan memangil fungsi sethash yang bertujuan untuk menghitung hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Timestamp: 0, Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	block.SetHash()
	return block
}
