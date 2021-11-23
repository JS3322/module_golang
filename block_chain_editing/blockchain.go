// package block_chain

// type Blockchain struct {
// 	blocks *[]Block
// }
// // new block add block chain
// func(bc *Blockchain) AddBlock(data string) {
// 	prevBlock   :=  bc.blocks[len(bc.blocks)-1]
// 	newBlock    :=  NewBlock(data, prevBlock.Hash)
// 	bc.blocks   :=  append(bc.blocks, newBlock)
//   }

// // instance block new bock chain
// func NewBlockchain() *Blockchain {
// 	return &Blockchain{[]*Block{NewGenesisBlock()}}
// }