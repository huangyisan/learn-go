package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// 转账需要考虑的为1. 转账数量 2. gas limit 3. gas price 4. 随机数nonce
	// 发送到网络之前,还需要发送方的私钥签名
	// 第三点gas price可以使用SuggestGasPrice方法从链上获取到合适的gas price
	// 第四点随机数nonce, 可以通过PendingNonceAt方法获取

	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}
	// 填写privatekey
	privateKey, err := crypto.HexToECDSA("6644cda75a855caa92a634ab1117339dc95ffe5dc10cd0c53b943a939f685c7d")
	if err != nil {
		log.Fatal(err)
	}
	// 这里是为了获取到nonce的值, 先从而获取到pulickkey,再从publickkey获取到钱包地址,也就是fromAddress, 然后拿到nonce
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 从publicECDSA中得到fromAddress
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 下面开始设置gas price , gas limit 等操作
	value := big.NewInt(1000000000000000000)                      // 单位为wei 这里为1个eth等待转账
	gasLimit := uint64(21000)                                     // gas limit
	gasPrice, err := client.SuggestGasPrice(context.Background()) // 获取gasprice的推荐值
	if err != nil {
		log.Fatal(err)
	}
	// 设定发送对端的钱包地址
	//toAddress := common.HexToHash("0x6E1Fb68b146B01ebB0d795Eea355449D705f4E23")
	address := common.HexToAddress("0x6E1Fb68b146B01ebB0d795Eea355449D705f4E23")
	// data数据跟合约交互,作为input的内容, 单纯转账一般不需要
	var data []byte
	LegacyTx := types.LegacyTx{
		Nonce:    nonce,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
		Value:    value,
		To:       &address,
	}
	// 组装交易信息
	tx := types.NewTx(&LegacyTx)

	// 获取ChainID后面签名交易信息需要使用
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易信息
	err = client.SendTransaction(context.Background(), signedTx)

	if err != nil {
		log.Fatal(err)
	}
	// 获取交易hash
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}
