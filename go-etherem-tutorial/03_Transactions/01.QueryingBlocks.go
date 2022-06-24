package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// number nil表示查询最后一个区块
	// 调用HeaderByNumber返回区块头信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// 打印区块id
	fmt.Println(header.Number.String())

	// 如果查询指定区块,则也是需要对其进行NewInt转换
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744 块id
	fmt.Println(block.Time())                // 1527211625 块生成时间戳
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065  块难度
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9 块的hash值
	fmt.Println(len(block.Transactions()))   // 144  块内交易数量

	// 可以使用TransactionCount来获取交易数量,传入blockhash
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
