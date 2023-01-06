package BLC

//交易输出管理

// 输出结构
type TxOutput struct {
	//金额
	Value        int
	ScriptPubKey string //UTXO的所有者
}
