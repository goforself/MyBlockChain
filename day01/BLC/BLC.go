package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

// 区块的基本结构与功能管理文件

// 实现一个基本的区块结构
type Block struct {
	TimeStamp     int64  //代表区块时间
	Hash          []byte //当前区块hash
	PrevBlockHash []byte //前一区块Hash
	Height        int64  //区块高度
	Data          []byte //存储交易数据（先不考虑merkle树）
}

// 新建区块
func NewBlock(prevBlockHash []byte, height int64, data []byte) *Block {
	var block Block
	block = Block{
		TimeStamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Data:          data,
	}
	block.SetHash()
	return &block
}

// 设置区块哈希
func (p *Block) SetHash() {
	//调用SAH256生成hash值
	//int64-》Hash
	timeStampBytes := IntToHex(p.TimeStamp)
	heightBytes := IntToHex(p.Height)
	//将多个[]byte数组转换为一个[]byte数组
	blockBytes := bytes.Join([][]byte{
		timeStampBytes,
		heightBytes,
		p.PrevBlockHash,
		p.Data,
	}, []byte{})
	hash := sha256.Sum256(blockBytes)
	p.Hash = hash[:] //将数组复制给切片
}

// 实现int64转为[]byte
func IntToHex(data int64) []byte {
	//建立缓冲流
	buffer := new(bytes.Buffer)
	//将data写入缓冲流
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("int transact to []byte failed!\n", err)
	}
	//将缓冲流以字节数组的方式读出
	return buffer.Bytes()
}
