package BLC

//交易输出管理

// 输出结构
type TxOutput struct {
	//金额
	Value        int
	ScriptPubKey string //UTXO的所有者
}

// 验证当前UTXO是否属于指定地址
func (txOut *TxOutput) CheckPubkeyWithAddress(address string) bool {
	return address == txOut.ScriptPubKey
}
