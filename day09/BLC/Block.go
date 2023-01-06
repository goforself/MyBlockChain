package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// 实现一个基本的区块结构
type Block struct {
	TimeStamp     int64          //代表区块时间
	Hash          []byte         //当前区块hash
	PrevBlockHash []byte         //前一区块Hash
	Height        int64          //区块高度
	Txs           []*Transaction //存储交易数据（交易列表，先不考虑merkle树）
	Nonce         int64          //运行POW算法时，生成的哈希变化值
}

// 新建区块
func NewBlock(prevBlockHash []byte, height int64, txs []*Transaction) *Block {
	var block Block
	block = Block{
		TimeStamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Txs:           txs,
	}
	//block.SetHash()
	//通过POW生成新的哈希
	pow := NewProofOfWork(&block)
	//执行工作量证明算法
	var hash []byte
	var nonce int64
	hash, nonce = pow.run()
	//fmt.Printf("Hash:%v\n", hash)
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

// 设置区块哈希
//func (p *Block) SetHash() {
//	//调用SAH256生成hash值
//	//int64-》Hash
//	timeStampBytes := IntToHex(p.TimeStamp)
//	heightBytes := IntToHex(p.Height)
//	//将多个[]byte数组转换为一个[]byte数组
//	blockBytes := bytes.Join([][]byte{
//		timeStampBytes,
//		heightBytes,
//		p.PrevBlockHash,
//		p.Data,
//	}, []byte{})
//	hash := sha256.Sum256(blockBytes)
//	p.Hash = hash[:] //将数组复制给切片
//}

// 生成创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(nil, 1, txs)
}

// 区块结构序列化，结构体->字节数组
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	//序列化结构体
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		log.Panicf("serialize the block to byte[] failed %v", err)
	}
	return buffer.Bytes()
}

// 区块结构反序列化，字节数组->结构体
func DeserializeBlock(blockByte []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockByte))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panicf("deserialize the byte[] to block failed %v", err)
	}
	return &block
}

// 交易结构txs序列化（实现方式类似于merkle树）
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	//将指定区块的交易哈希进行拼接
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
