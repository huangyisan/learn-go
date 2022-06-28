package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	// 生成签名的组件是:
	// 1. 签名者私钥  2.待前面数据的哈希

	// 加载私钥
	privateKey, err := crypto.HexToECDSA("6644cda75a855caa92a634ab1117339dc95ffe5dc10cd0c53b943a939f685c7d")
	if err != nil {
		log.Fatal(err)
	}
	//签名数据假设是hello, 使用Keccak256Hash方法先对其进行hash操作
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	// 最后使用私钥对其签名hash, 得到签名
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(signature))
	fmt.Println(signature)
}
