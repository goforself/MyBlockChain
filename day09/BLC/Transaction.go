package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 交易管理文件
type Transaction struct {
	TxHash []byte      //当前区块交易哈希
	Vins   []*TxInput  //输入列表
	Vouts  []*TxOutput //输出列表
}

// 实现coinbase交易
func NewCoinbaseTransaction(address string) *Transaction {
	var txCoinbase *Transaction
	//txCoinbase的特点
	//输入
	//txHash：nil
	//Vout：-1
	//ScriptSig：system reward
	txInput := &TxInput{TxHash: nil, Vout: -1, ScriptSig: "system reward"}
	//输出
	//value:10
	//address:
	txOutput := &TxOutput{value: 10, ScriptPubKey: address}
	txCoinbase = &Transaction{TxHash: nil,
		Vins:  []*TxInput{txInput},
		Vouts: []*TxOutput{txOutput}}
	//交易哈希生成
	txCoinbase.GenerateTransactionHash()
	return txCoinbase
}

// 实现交易的序列化
func (tx *Transaction) GenerateTransactionHash() {
	var res bytes.Buffer
	//设置编码对象
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panicf("computer the Transaction hash failed%v\n", err)
	}
	//生成哈希值
	hash := sha256.Sum256(res.Bytes())
	tx.TxHash = hash[:]
}
