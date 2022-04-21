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

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

func main() {

	client, err := ethclient.Dial("wss://bsc-ws-node.nariox.org")

	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x26193C7fa4354AE49eC53eA2cEBC513dc39A10aa")
	//add := common.HexToAddress("0xF489d5EfA3dA8B426F86cb4Df30edBf3f321913d")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(17104780),
		ToBlock:   big.NewInt(17104790),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("abi.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	contractAbi, err := abi.JSON(strings.NewReader(string(content)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				fmt.Println("xxxxx")
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Values: %s\n", transferEvent.Value.String())
		default:
			fmt.Println("no log")
		}
	}
}
