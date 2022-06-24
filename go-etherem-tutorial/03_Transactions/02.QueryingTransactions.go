package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(2)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// block里面用Transactions方法来读取块中的事务,会返回一个Transaction类型的列表.
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 交易hash 0xb14a1286ce38c8c22020e1df7c6fbe194156c3341b77e6b88557f189b7099d46
		fmt.Println(tx.Value().String())    // 1000000000000000000 . 交易数量
		fmt.Println(tx.Gas())               // 21000 . gas limit 的值
		fmt.Println(tx.GasPrice().Uint64()) // 20000000000  总共的gas费用
		fmt.Println(tx.Nonce())             // 0  Nonce唯一值
		fmt.Println(tx.Data())              // [] 交易中input的数据
		fmt.Println(tx.To().Hex())          // 接受者的地址 0x6E1Fb68b146B01ebB0d795Eea355449D705f4E23

		// 通过NetworkID方法 获取chain id
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		baseFee := big.NewInt(0)
		// 查看发送者的钱包地址, 使用AsMessage方法
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), baseFee); err == nil {
			fmt.Println(chainID)
			fmt.Println(msg.From().Hex())
		} else {
			log.Fatal(err)
		}
		// 使用TransactionReceipt方法查看交易结果,成功还是失败, 1 为成功 0 为失败
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)

		// 可以通过块hash配合TransactionInBlock方法来查看,块中有多少个交易事务,并且获取对应的交易事务
		// 0x51815fa429ee2bf85566500e36d31490097dcfd67d4286b9839905ed3725b969为块hash
		blockHash := common.HexToHash("0x51815fa429ee2bf85566500e36d31490097dcfd67d4286b9839905ed3725b969")
		count, err := client.TransactionCount(context.Background(), blockHash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(count)
		for idx := uint(0); idx < count; idx++ {
			tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
			if err != nil {
				log.Fatal(err)
			}
			// 打印出交易hash 0xc885507d45c16dc20dd5a99659571c20d277f4261076df628bb06a0e1b70edb6
			fmt.Println(tx.Hash().Hex())
		}

		// 使用TransactionByHash 查询给定交易hash的具体单个事务
		txHash := common.HexToHash("0xc885507d45c16dc20dd5a99659571c20d277f4261076df628bb06a0e1b70edb6")
		tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Value().String())    // 1000000000000000000 . 交易数量
		fmt.Println(tx.Gas())               // 21000 . gas limit 的值
		fmt.Println(tx.GasPrice().Uint64()) // 20000000000  总共的gas费用
		fmt.Println(tx.Nonce())             // 0  Nonce唯一值
		fmt.Println(tx.Data())              // [] 交易中input的数据
		fmt.Println(tx.To().Hex())          // 接受者的地址 0x6E1Fb68b146B01ebB0d795Eea355449D705f4E23

		fmt.Println(tx.Hash().Hex()) // 0xc885507d45c16dc20dd5a99659571c20d277f4261076df628bb06a0e1b70edb6
		fmt.Println(isPending)       // false  查看交易是否正在处理中
	}
}
