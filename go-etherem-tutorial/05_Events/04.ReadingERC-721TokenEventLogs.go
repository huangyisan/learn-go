package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// LogTransfer ..
type LogTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

func main() {
	client, err := ethclient.Dial("http://54.178.71.35/rpc")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x416f1d70c1c22608814d9f36c492efb3ba8cad4c")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(17104780),
		ToBlock:   big.NewInt(17104790),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//file, err := os.Open("fish_abi.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//content, err := ioutil.ReadAll(file)
	//contractAbi, err := abi.JSON(strings.NewReader(string(content)))
	//contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	//LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	//logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTokenTransfer

			//err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			//if err != nil {
			//	fmt.Println("error")
			//	log.Fatal(err)
			//}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			// tokens 处于topics中第三个位置
			fmt.Printf("Tokens: %s\n", vLog.Topics[3].Big())
			//fmt.Printf("Tokens: %v\n", transferEvent)

			fmt.Printf("\n\n")
		}
	}
}
