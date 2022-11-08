package main

import (
	"MyBlockChain/day01/BLC"
	"fmt"
)

// 启动
func main() {
	block := BLC.NewBlock([]byte{1}, 15, []byte("the first block"))
	block.SetHash()
	fmt.Printf("the first block :%v\n", block)
}
