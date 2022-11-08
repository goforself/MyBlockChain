package BLC

//区块链管理文件

// 区块链基本结构
type BlockChain struct {
	Blocks []*Block //区块链的切片
}

// 添加区块到区块链中
func (bc *BlockChain) AddBlock(prevBlockHash []byte, height int64, data []byte) {
	var newBlock *Block
	newBlock = NewBlock(prevBlockHash, height, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//生成创世区块
	block := CreateGenesisBlock([]byte("the first block"))
	return &BlockChain{[]*Block{block}}
}
