package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//区块链管理文件

// 数据库名称
const dbName = "block.db"

// 表名称
const blockTableName = "blocks"

// 区块链基本结构
type BlockChain struct {
	//Blocks []*Block //区块链的切片
	DB  *bolt.DB //以存储区块链到数据库中
	Tip []byte   //保存最新区块的哈希值
}

// 添加区块到区块链中
func (bc *BlockChain) AddBlock(data []byte) {
	//更新区块数据（insert）
	bc.DB.Update(func(tx *bolt.Tx) error {
		//获取数据库桶
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//获取对应区块
			blockTypes := b.Get(bc.Tip)
			//区块结构反序列化
			latestBlock := DeserializeBlock(blockTypes)
			//新建区块,传入参数：prevBlockHash []byte, height int64, data []byte
			newBlock := NewBlock(latestBlock.Hash, latestBlock.Height+1, data)
			//存入数据库
			bc.Tip = newBlock.Hash
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panicf("insert the new block failed %v\n", err)
			}
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panicf("save the latest block hash failed %v\n", err)
			}
		}
		return nil
	})
}

// 初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//存储最新区块链哈希
	var blockHash []byte
	//1.创建或打开一个数据库
	db, err := bolt.Open(dbName, 0600, nil)
	//2.创建一个桶,将创世区块存入数据库中
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//数据表不存在
			b, err := tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panicf("create backet [%s] failed %v\n", blockTableName, err)
			}
			//生成创世区块
			genesisBlock := CreateGenesisBlock([]byte("the first block"))
			//键为区块的哈希，值为区块的序列化
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panicf("insert the genesis block failed %v\n", err)
			}
			blockHash = genesisBlock.Hash
			//数据库中也存储最新区块的哈希
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panicf("save the latest hash of genesis block failed %v\n", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panicf("create db [%s] failed %v\n", dbName, err)
	}
	return &BlockChain{DB: db, Tip: blockHash}
}
