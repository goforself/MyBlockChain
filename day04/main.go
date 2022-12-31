package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		//创建一个桶
		b, err := tx.CreateBucket([]byte("MyBucket2"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		//写入数据
		if b != nil {
			err := b.Put([]byte("11"), []byte("abc"))
			if err != nil {
				return err
			}
		}
		return nil
	})
	//读出数据
	db.View(func(tx *bolt.Tx) error {
		//获取桶
		b := tx.Bucket([]byte("MyBucket2"))
		if nil != b {
			value := b.Get([]byte("11"))
			fmt.Printf("value:%s\n", value)

		}
		return nil
	})
}
