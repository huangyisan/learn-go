package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 订阅事件则用wss协议去连接node节点,然后订阅合约地址
	client, err := ethclient.Dial("https://u0lvw6c3lk:7vUkST4yxgKT_yjF5laCagbtYNCkw0ZOQIYv-FteCvE@u0avnfe4vh-u0j18z3adv-rpc.us0-aws.kaleido.io/")
	//client, err := ethclient.Dial("ws://localhost:7545")
	//rpcClient, err := rpc.Dial("wss://u0avnfe4vh-u0j18z3adv-wss.us0-aws.kaleido.io/")

	if err != nil {
		log.Fatal(err)
	}

	// 创建筛选查询
	contractAddress := common.HexToAddress("0x9c7cad41c2a2b9aefe3721386743321f098cc6cc")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 接受时间的方式是通过Go channel. 一个types.Log的channel
	logs := make(chan types.Log)

	// 客户端调用SubscribeFilterLogs来订阅.
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// 使用select语句设置一个连续循环来读入新的日志事件或订阅错误
	fmt.Println("start to subscribe")
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}
