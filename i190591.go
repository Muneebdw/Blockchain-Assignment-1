// Language: go
// Path: i190593.go
package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Blockchain struct {
	// array of pointers, pointing to different blocks
	blocks []*Block
}

type Block struct {
	data      string
	nonce     int
	chash     []byte
	prevhashs []byte

	//create a createhash function

}

// createhash
func (b *Block) createhash() []byte {
	//create a string to store the data
	info := bytes.Join([][]byte{b.prevhashs, []byte(b.data), []byte(strconv.Itoa(b.nonce))}, []byte{})
	//create a hash
	hash := sha256.Sum256(info)
	return hash[:]
}

/*
func NewBlock(transaction string, nonce int, previousHash string) *block {
A method to add new block. To keep things simple, you could provide a
sting of your choice as a transaction (e.g., “bob to alice”). Also, use
any integer value as a nonce. The CreateHash() method will provide you the
block Hash value.
*/
func NewBlock(data string, nonce int, prevhashs []byte) *Block {
	block := &Block{data, nonce, []byte{}, prevhashs}
	//create a hash
	block.chash = block.createhash()
	return block
}

/*
func DisplayBlocks() {
A method to print all the blocks in a nice format showing block data such
as transaction, nonce, previous hash, current block hash
*/
func (Blockchains *Blockchain) DisplayBlocks() {
	for _, block := range Blockchains.blocks {
		fmt.Printf("Data: %s", block.data)
		fmt.Printf("Nonce: %d", block.nonce)
		fmt.Printf("Previous Hash: %x", block.prevhashs)
		fmt.Printf("Current Hash: %x", block.chash)
		fmt.Println()
	}
}

/*
func ChangeBlock() {
function to change block transaction of the given block ref
*/
func (Blockchains *Blockchain) ChangeBlock(block *Block, data string) {
	block.data = data
	block.chash = block.createhash()
}

/*
func VerifyChain() {
function to verify blockchain in case any changes are made.
*/
func (Blockchains *Blockchain) VerifyChain() bool {
	for i := 1; i < len(Blockchains.blocks); i++ {
		prevBlock := Blockchains.blocks[i-1]
		currentBlock := Blockchains.blocks[i]
		if bytes.Compare(currentBlock.prevhashs, prevBlock.chash) != 0 {
			return false
		}
		if bytes.Compare(currentBlock.chash, currentBlock.createhash()) != 0 {
			return false
		}
	}
	return true
}

/*
func CalculateHash (stringToHash string) {
function for calculating hash of a block
*/
func CalculateHash(stringToHash string) []byte {
	hash := sha256.Sum256([]byte(stringToHash))
	return hash[:]
}

func main() {
	//create a blockchain
	Blockchains := Blockchain{}
	//create a genesis block
	genesisBlock := NewBlock("Genesis Block", 0, []byte{})
	//add the genesis block to the blockchain
	Blockchains.blocks = append(Blockchains.blocks, genesisBlock)
	//create a new block
	newBlock := NewBlock("New Block", 0, genesisBlock.chash)
	//add the new block to the blockchain
	Blockchains.blocks = append(Blockchains.blocks, newBlock)
	//display the blocks
	Blockchains.DisplayBlocks()
	//change the data of the new block
	Blockchains.ChangeBlock(newBlock, "Changed Block")
	//display the blocks
	Blockchains.DisplayBlocks()
	//verify the blockchain
	fmt.Println(Blockchains.VerifyChain())

}
