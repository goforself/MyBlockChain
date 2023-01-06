package BLC

// UTXO结构管理
type UTXO struct {
	//交易区块哈希
	txHash []byte
	//Vout的索引
	index int
	//Vout的值
	output *TxOutput
}
