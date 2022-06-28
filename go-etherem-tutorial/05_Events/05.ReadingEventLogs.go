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

// LogTransfer ..
type FeeBurned struct {
	Amount *big.Int
}

func main() {
	client, err := ethclient.Dial("https://bsc-dataseed3.binance.org")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x0000000000000000000000000000000000001000")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(19055442),
		ToBlock:   big.NewInt(19055442),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("validate_abi.json")
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

	feeBurnedSig := []byte("feeBurned(uint256)")
	//LogApprovalSig := []byte("Approval(address,address,uint256)")
	feeBurnedSigHash := crypto.Keccak256Hash(feeBurnedSig)
	//logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case feeBurnedSigHash.Hex():
			fmt.Printf("Log Name: feeBurnedSig\n")

			var feeBurned FeeBurned
			// indexed的数据是不需要unpack的,唯有当需要读取非indexed的数据才需要.
			err := contractAbi.UnpackIntoInterface(&feeBurned, "feeBurned", vLog.Data)
			if err != nil {
				fmt.Println("error")
				log.Fatal(err)
			}

			//transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			//transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			//fmt.Printf("From: %s\n", transferEvent.From.Hex())
			//fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("feeBurned: %s\n", feeBurned.Amount.String())
			//fmt.Printf("Tokens: %v\n", transferEvent)

			fmt.Printf("\n\n")
		}
	}
}
