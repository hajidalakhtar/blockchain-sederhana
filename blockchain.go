package main

// membuat struct blockchain
type Blockchain struct {
	blocks []*Block
}

// membuat fungsi untuk membuat blockchain baru dengan block kosong awal
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// membuat fungsi untuk menambahkan block baru ke dalam rantai block (blockchain)
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]   // mengambil block terakhir dari blockchain
	newBlock := NewBlock(data, prevBlock.Hash) // membuat block baru dengan memanggil fungsi NewBlock dengan mengirim data dan hash sebelumnya
	bc.blocks = append(bc.blocks, newBlock)    // menambahkan block baru ke dalam blockchain
}

// dikarenakan tidak ada block di dalam blockchain, function ini bertujuan untuk membuat block kosong awal
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // memanggil fungsi NewBlock dengan mengirim data dan hash kosong
}
