package BLC

import (
	"fmt"
	"os"
)

// 打印区块完整信息
func (cli *CLI) printChain() {
	if !DBExist() {
		fmt.Printf("DB haven't existed")
		os.Exit(1)
	}
	block := BlockChainObject()
	block.PrintBlockChain()
}
