package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// connect to ganache
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client

	client1, err := ethclient.Dial("https://u0lvw6c3lk:7vUkST4yxgKT_yjF5laCagbtYNCkw0ZOQIYv-FteCvE@u0avnfe4vh-u0j18z3adv-rpc.us0-aws.kaleido.io/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client1

}
