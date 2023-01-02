package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//区块链迭代器管理文件

// 迭代器基本结构
type BlockChainIterator struct {
	DB      *bolt.DB //迭代目标
	curHash []byte   //当前迭代目标
}

// 创建迭代器对象
func (blc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{DB: blc.DB, curHash: blc.Tip}
}

// 实现迭代函数next（）
func (iter *BlockChainIterator) Next() *Block {
	var curBlock *Block
	err := iter.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			Block := b.Get(iter.curHash)
			curBlock = DeserializeBlock(Block)
			//更新迭代器当前指向结点哈希
			iter.curHash = curBlock.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panicf("Iterator the db failed%v\n", err)
	} else {
		return curBlock
	}
	return nil
}
