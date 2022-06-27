package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("https://bsc-dataseed3.binance.org")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x26193C7fa4354AE49eC53eA2cEBC513dc39A10aa")
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

	file, err := os.Open("abi.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	contractAbi, err := abi.JSON(strings.NewReader(string(content)))
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

			//case logApprovalSigHash.Hex():
			//	fmt.Printf("Log Name: Approval\n")
			//
			//	var approvalEvent LogApproval
			//
			//	err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			//	if err != nil {
			//		log.Fatal(err)
			//	}
			//
			//	approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			//	approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
			//
			//	fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			//	fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			//	fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
			//}

			fmt.Printf("\n\n")
		}
	}
}
