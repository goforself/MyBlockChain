package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// 对blockchain进行command lina的管理
type CLI struct {
	BC *BlockChain
}

// 命令行提示用法展示
func PrintUsage() {
	fmt.Println("Usage:  go run file.exe <command> [arguments]")
	//初始化区块链
	fmt.Printf("\tcreate    	-address         create a new blockchain\n")
	//添加区块
	fmt.Printf("\taddBlock   -data  	         add a new block\n")
	//遍历区块链
	fmt.Printf("\tprintChain                  print the all message of blockchain\n")
}

// 初始化区块链
func (cli *CLI) createBlockchain(address string) {
	CreateBlockChainWithGenesisBlock(address)
}

// 添加区块
func (cli *CLI) addBlock(data []*Transaction) {
	if !DBExist() {
		fmt.Printf("DB haven't existed")
		os.Exit(1)
	}
	block := BlockChainObject()
	block.AddBlock(data)
}

// 打印区块完整信息
func (cli *CLI) printChain() {
	if !DBExist() {
		fmt.Printf("DB haven't existed")
		os.Exit(1)
	}
	block := BlockChainObject()
	block.PrintBlockChain()
}

// 参数数量检测函数
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}
}

// 命令行运行函数
func (cli *CLI) Run() {
	IsValidArgs()
	//命令设置
	//添加区块
	addBlockcmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	//输出区块完整信息
	printChaincmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	//创建区块链
	createBLCwithGenesisBlockcmd := flag.NewFlagSet("create", flag.ExitOnError)

	//参数默认值设置
	flagAddBlockArg := addBlockcmd.String("data", "Alice send 100 eth to Bob", "add data to block")
	flagCreateGenesisBlockArg := createBLCwithGenesisBlockcmd.String("address", "czd", "the first system reward address")

	//判断命令
	switch os.Args[1] {
	case "addBlock":
		err := addBlockcmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse the addBlockcmd failed%v\n", err)
		}
	case "printChain":
		err := printChaincmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse the printChaincmd failed%v\n", err)
		}
	case "create":
		err := createBLCwithGenesisBlockcmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse the createBLCwithGenesisBlockcmd failed%v\n", err)
		}
	default:
		//没有传递任何命令或不在命令列表中
		PrintUsage()
		os.Exit(1)
	}

	//添加区块
	if addBlockcmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock([]*Transaction{})
	}
	//打印区块链信息
	if printChaincmd.Parsed() {
		fmt.Printf("\nactivate\n")
		cli.printChain()
	}
	//创建区块链
	if createBLCwithGenesisBlockcmd.Parsed() {
		if *flagCreateGenesisBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockchain(*flagCreateGenesisBlockArg)
	}
}
