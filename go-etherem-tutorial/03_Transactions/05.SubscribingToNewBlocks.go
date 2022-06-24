package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 订阅websocket节点
	client, err := ethclient.Dial("wss://infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个接受最新区块的channel
	headers := make(chan *types.Header)

	// 使用SubscribeNewHead方法,返回一个订阅对象
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 用select来监听订阅.
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			// 要获得该区块的完整内容, 我们可以使用BlockByHash方法
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        //
			fmt.Println(block.Number().Uint64())   //
			fmt.Println(block.Time())              //
			fmt.Println(block.Nonce())             //
			fmt.Println(len(block.Transactions())) //
		}

	}

}
