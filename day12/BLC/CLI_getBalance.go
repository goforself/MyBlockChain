package BLC

import "fmt"

// 查询余额
func (cli *CLI) getBalance(from string) {
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	balance := blockchain.getBalance(from)
	fmt.Printf("tha balance of [%s] : %d\n", from, balance)
}
