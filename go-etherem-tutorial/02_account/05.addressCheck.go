package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func main() {
	// 合法钱包地址和合约地址是以0x开头后面跟随40位16进制
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	// 区别钱包地址还是合约地址的方法是查看其字节码长度,如果大于0,则是合约地址,反之是钱包地址
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	// 使用CodeAt方法获取bytecode
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract)

	// 这是一个正常的钱包地址
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
