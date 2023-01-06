package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//共识算法管理文件

//实现POW实例以及相关功能

// 目标难度值,前targetBit位为0
const targetBit = 16

// 工作量证明的结构
type ProofOfWork struct {
	//需要共识验证的区块
	Block *Block
	//目标难度的哈希,大数存储
	target *big.Int
}

// 创建一个POW对象
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//Lsh函数为：target向左移动（256-targetBit）位
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{Block: block, target: target}
}

// 执行POW算法，比较哈希
// 返回哈希与碰撞次数
func (pow *ProofOfWork) run() ([]byte, int64) {
	//碰撞次数
	var nonce = int64(0)
	//用于比较的大数hash
	var hashInt big.Int
	//目标hash值
	var hash [32]byte
	//无限循环，生成符合条件的哈希
	for {
		//生成准备数据
		dataByte := pow.prepareData(int64(nonce))
		hash = sha256.Sum256(dataByte)
		//将byte数据转换为大数
		hashInt.SetBytes(hash[:])
		if pow.target.Cmp(&hashInt) == 1 {
			//找到了符合条件的hash
			break
		}
		nonce++
	}
	fmt.Printf("打印碰撞次数%v\n", nonce)
	return hash[:], nonce
}

// 生成准备数据，对ProofOfWork数据拼接形成哈希值并返回
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	var data []byte
	timeStampBytes := IntToHex(pow.Block.TimeStamp)
	heightBytes := IntToHex(pow.Block.Height)
	//将多个[]byte数组转换为一个[]byte数组
	data = bytes.Join([][]byte{
		timeStampBytes,
		heightBytes,
		pow.Block.PrevBlockHash,
		pow.Block.HashTransactions(),
		IntToHex(nonce),
		IntToHex(targetBit),
	}, []byte{})
	return data
}
