package main

import (
	"PublicChain/chain"
	"fmt"
)

func main() {
	fmt.Println("hellow world")
	genesis :=chain.CreateGenesisBlock([]byte("hello world"))
	fmt.Println("新区快0",genesis)
	block1 :=chain.CreateBlock(genesis.Height,genesis.Hash,nil)
	fmt.Println("新区快1",block1)

}
