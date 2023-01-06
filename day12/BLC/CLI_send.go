package BLC

import (
	"fmt"
	"os"
)

// 发起交易
func (cli *CLI) send(from, to, amount []string) {
	if !DBExist() {
		fmt.Printf("DB haven't existed")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	//挖矿
	//windows下输入格式要改变： .\bc.exe send -from '[\"Tom\"]' -to '[\"Alice\"]' -amount '[\"3\"]'
	blockchain.MineNewBlock(from, to, amount)
}
