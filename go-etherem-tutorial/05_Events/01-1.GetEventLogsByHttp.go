package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	// 订阅事件则用wss协议去连接node节点,然后订阅合约地址
	client, err := ethclient.Dial("https://u0lvw6c3lk:7vUkST4yxgKT_yjF5laCagbtYNCkw0ZOQIYv-FteCvE@u0avnfe4vh-u0j18z3adv-rpc.us0-aws.kaleido.io/")
	//client, err := ethclient.Dial("ws://localhost:7545")
	//rpcClient, err := rpc.Dial("wss://u0avnfe4vh-u0j18z3adv-wss.us0-aws.kaleido.io/")

	if err != nil {
		log.Fatal(err)
	}

	// 合约中Transfer方法结构体
	type LogTransfer struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	// 创建筛选查询
	contractAddress := common.HexToAddress("0x9c7cad41c2a2b9aefe3721386743321f098cc6cc")
	// block range 51150 - 51160
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(51150),
		//ToBlock:   big.NewInt(51160),
		Addresses: []common.Address{contractAddress},
	}

	// 客户端调用SubscribeFilterLogs来订阅.
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 使用select语句设置一个连续循环来读入新的日志事件或订阅错误
	fmt.Println("start to subscribe")
	// 函数签名
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	// 读取abi文件
	file, err := os.Open("kaleido_abi.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	contractAbi, err := abi.JSON(strings.NewReader(string(content)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")
			var transferEvent LogTransfer
			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				fmt.Println("error")
				log.Fatal(err)
			}
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Value.String())
		}
	}
}
