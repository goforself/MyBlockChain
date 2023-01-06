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
	fmt.Printf("\tcreate        -address                                  create a new blockchain\n")
	//添加区块
	fmt.Printf("\taddBlock      -data                                     add a new block\n")
	//遍历区块链
	fmt.Printf("\tprintChain                                              print the all message of blockchain\n")
	//添加命令行转账
	fmt.Printf("\tsend          -from  FROM   -to  TO   -amount  AMOUNT   Generate transaction\n")
	fmt.Printf("\t\tthe send of usage:\n")
	fmt.Printf("\t\t-from   FROM    --the source address of transaction\n")
	fmt.Printf("\t\t-to     TO      --the destination address of transaction\n")
	fmt.Printf("\t\t-amount AMOUNT  --the value of transaction\n")
	//余额查询
	fmt.Printf("\tgetBalance    -address                                  get the balance of address ")
	fmt.Printf("\t\tthe getBalance of usage:\n")
	fmt.Println("\t\t-address the address that you want to get balance of")
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

// 发起交易
func (cli *CLI) send(from, to, amount []string) {
	if !DBExist() {
		fmt.Printf("DB haven't existed")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	//挖矿
	//windows下输入格式要改变： .\bc.exe send -from '[\"Tom\"]' -to '[\"Alice\"]' -amount '[\"3\"]'
	blockchain.MineNewBlock(from, to, amount)
}

// 查询余额
func (cli *CLI) getBalance(from string) {
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	balance := blockchain.getBalance(from)
	fmt.Printf("tha balance of [%s] : %d\n", from, balance)
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
	//产生交易
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	//查询余额
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)

	//参数默认值设置
	flagAddBlockArg := addBlockcmd.String("data", "Alice send 100 eth to Bob", "add data to block")
	flagCreateGenesisBlockArg := createBLCwithGenesisBlockcmd.String("address", "czd", "the first system reward address")
	flagSendFromArg := sendCmd.String("from", "", "the source address of transaction")
	flagSendToArg := sendCmd.String("to", "", "the destination address of transaction")
	flagSendAmountArg := sendCmd.String("amount", "", "the value of transaction")
	flagGetBalanceArg := getBalanceCmd.String("address", "", "the balance of address")
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
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse the send failed%v\n", err)
		}
	case "getBalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicf("parse the getBalance failed%v\n", err)
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
	if sendCmd.Parsed() {
		if *flagSendFromArg == "" {
			fmt.Printf("-from : the arg of from can't be nil\n")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendToArg == "" {
			fmt.Printf("-to : the arg of to can't be nil\n")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendAmountArg == "" {
			fmt.Printf("-amount : the arg of amount can't be nil\n")
			PrintUsage()
			os.Exit(1)
		}
		cli.send(JSONToSlice(*flagSendFromArg), JSONToSlice(*flagSendToArg), JSONToSlice(*flagSendAmountArg))
		fmt.Printf("\tFROM:[%s]\n", JSONToSlice(*flagSendFromArg))
		fmt.Printf("\tTO:[%s]\n", JSONToSlice(*flagSendToArg))
		fmt.Printf("\tAMOUNT:[%s]\n", JSONToSlice(*flagSendAmountArg))
	}
	if getBalanceCmd.Parsed() {
		if *flagGetBalanceArg == "" {
			fmt.Printf("-address : the arg of address can't be nil\n")
			PrintUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceArg)
	}
}
