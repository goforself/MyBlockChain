package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
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
func (bc *BlockChain) AddBlock(data []*Transaction) {
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

// 判断数据库文件是否存在
func DBExist() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		//数据库文件不存在
		return false
	}
	return true
}

// 初始化区块链
func CreateBlockChainWithGenesisBlock(address string) *BlockChain {
	if DBExist() {
		fmt.Println("the genesis bolck existed")
		os.Exit(1)
	}
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
			//生成一个coinbase交易
			txCoinbase := NewCoinbaseTransaction(address)
			//生成创世区块
			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})
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
		} else {
			//数据表已存在
			blockHash = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panicf("create db [%s] failed %v\n", dbName, err)
	}
	return &BlockChain{DB: db, Tip: blockHash}
}

// 遍历区块链，输出所有区块信息
func (bc *BlockChain) PrintBlockChain() {
	var curBlock *Block
	var iter *BlockChainIterator = bc.Iterator()
	//读取数据库
	fmt.Printf("打印区块链完整信息。。。\n")
	//循环读取
	for {
		curBlock = iter.Next()
		fmt.Printf("-------------------------------------------\n")
		fmt.Printf("\tHash : %x\n", curBlock.Hash)
		fmt.Printf("\tPrevBlockHash : %x\n", curBlock.PrevBlockHash)
		fmt.Printf("\tHeight : %x\n", curBlock.Height)
		fmt.Printf("\tNonce : %x\n", curBlock.Nonce)
		fmt.Printf("\tTimeStamp : %x\n", curBlock.TimeStamp)
		fmt.Printf("\tTxs : \n")
		for _, tx := range curBlock.Txs {
			//交易哈希
			fmt.Printf("\t\ttx-Hash  : %x\n", tx.TxHash)
			fmt.Printf("\t\tinput:\n")
			for index, vin := range tx.Vins {
				fmt.Printf("\t\t\tinput-index:%d\n", index)
				//上一个交易哈希
				fmt.Printf("\t\t\ttx-Vins  : %x\n", vin.TxHash)
				//上一个交易索引
				fmt.Printf("\t\t\ttx-Vout  : %x\n", vin.Vout)
				//上一个交易签名
				fmt.Printf("\t\t\ttx-scriptSig  : %s\n", vin.ScriptSig)
			}
			fmt.Printf("\t\toutput:\n")
			for index, vout := range tx.Vouts {
				fmt.Printf("\t\t\toutput-index:%d\n", index)
				//转账金额
				fmt.Printf("\t\t\ttx-value  : %x\n", vout.Value)
				//转账对象
				fmt.Printf("\t\t\ttx-scriptPubKey  : %s\n", vout.ScriptPubKey)
			}
		}
		//退出条件
		var hashInt big.Int
		hashInt.SetBytes(iter.curHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			//遍历到了创世区块
			break
		}

	}
}

// 从db里面获取一个blockchain对象
func BlockChainObject() *BlockChain {
	var latestHash []byte
	//获取DB
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicf("the database haven't exist")
	}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		//获取tip
		latestHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panicf("get the blockchain object failed%v\n", err)
	}
	return &BlockChain{db, latestHash}
}

// 实现挖矿功能
// 通过接受交易，生成区块
func (bc *BlockChain) MineNewBlock(from []string, to []string, amount []string) {
	//搁置交易生成步骤
	var txs []*Transaction
	var block *Block
	//字符串转数字
	value, _ := strconv.Atoi(amount[0])
	//产生新的交易
	tx := NewSampleTransaction(from[0], to[0], value)
	txs = append(txs, tx)
	bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//得到最新区块哈希
			latestHash := b.Get([]byte("l"))
			//得到最新区块
			blockByte := b.Get(latestHash)
			//反序列化最新区块
			block = DeserializeBlock(blockByte)
		}
		return nil
	})
	//挖掘区块
	block = NewBlock(block.Hash, block.Height+1, txs)
	bc.DB.Update(func(tx *bolt.Tx) error {
		//更新数据库的最新区块信息
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			b.Put([]byte("l"), block.Hash)
			b.Put(block.Hash, block.Serialize())
		}
		return nil
	})
}
