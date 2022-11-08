package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

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
