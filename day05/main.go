package main

import "MyBlockChain/day05/BLC"

// 启动
func main() {
	//生成创世区块
	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	//产生第一个区块
	blockChain.AddBlock([]byte("Alice send 100 eth to Bob"))
	//产生第二个区块
	blockChain.AddBlock([]byte("Bob send 150 eth to Tom"))
}
