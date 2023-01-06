package BLC

//交易输出管理

// 输出结构
type TxOutput struct {
	//金额
	value        int
	ScriptPubKey string //UTXO的所有者
}
