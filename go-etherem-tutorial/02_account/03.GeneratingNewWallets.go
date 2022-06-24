package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

// 生成钱包其实是生成私钥, 然后将公钥进行Keccak-256运算,就能得到钱包地址.

func main() {
	// 使用generatekey生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// 将得到的转换为字节,在用hexutil包的Encode方法转换成16进制字符串,剔除前面0x两位就得到了人能看的私钥了.
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 这里得到的就是私钥了.
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	// 通过私钥可以返回公钥
	publicKey := privateKey.Public()

	// 然后在将其转换为16进制.然后Encode处理, 去掉前面四个字符,0x04, 这个是EC前缀
	// 这里断言判断下是否为publickey
	publickeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("not public key")
	}
	publicKeyBytes := crypto.FromECDSAPub(publickeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// 有了公钥就能通过公钥看到公共地址,也就是钱包地址,用Keccak256方法
	hash := sha3.NewLegacyKeccak256()
	//fmt.Println(publicKeyBytes)
	fmt.Println(publicKeyBytes[1:])
	// 去掉第一个字节
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)))
	// 取最后40个字符(20字节),总共64个字符(32字节),
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

}
