package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x112Bff1843d6016b1481876b65d694c678718454")
	// 区块号必须是big.Int类型
	blockNumber := big.NewInt(0)
	// 使用BalanceAt来获取账户余额
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。

	// 读取wei的值.必须做计算 wei/10^18
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	// fbalance除以big.NewFloat(math.Pow10(18)) 四舍五入的商
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	// 使用PendingBalanceAt查询执行处理中,或等待处理中期间,钱包的余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
