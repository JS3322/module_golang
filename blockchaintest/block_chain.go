// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// 새로운 블록체인 생성
// 	bc := NewBlockchain()

// 	// 새로운 블록들 생성
// 	bc.AddBlock("Send 1 BTC to Ivan")
// 	bc.AddBlock("Send 2 more BTC to Ivan")
// 	// 블록체인 출력
// 	for _, block := range bc.blocks {
// 		fmt.Printf("Prev : %x\n", block.PrevBlockHash)
// 		fmt.Printf("Data : %s\n", block.Data)
// 		fmt.Printf("Hash : %x\n", block.Hash)
// 		fmt.Println()
// 	}
// }