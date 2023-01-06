package BLC

//交易输入管理

// 输入结构
type TxInput struct {
	TxHash    []byte //交易哈希(不是当前区块交易哈希)
	Vout      int    //引用上一笔交易的输出索引号
	ScriptSig string //用户名
}
