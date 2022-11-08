package main

import (
	"MyBlockChain/day02/BLC"
	"fmt"
)

// 启动
func main() {
	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	//fmt.Printf("the first block :%v\n", blockChain)
	blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Hash,
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, []byte("Alice send 10 btc to bob"))
	blockChain.AddBlock(blockChain.Blocks[len(blockChain.Blocks)-1].Hash,
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, []byte("Bob send 5 btc to Tom"))
	for _, d := range blockChain.Blocks {
		fmt.Printf("prevHash:%v\nnowHash:%v\n", d.PrevBlockHash, d.Hash)
	}
}
