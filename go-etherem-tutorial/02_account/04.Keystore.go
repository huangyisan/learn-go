package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func createKs() {
	// 使用NewKeyStore在本地生成一个keystore,是一个以文件方式,用来存储密钥对的文件.需要传入密码才可以.
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

// 导入keystore, 用Import方法,该方法接收keystore的JSON数据作为字节。第二个参数是用于加密私钥的口令。第三个参数是指定一个新的加密口令
// 这个导入其实更像是对原始的账户修改密码后写入本地新目录
func importKs() {
	file := "./wallet/UTC--2022-06-24T04-43-35.254549000Z--382277f70c21fb360ddc8bad5f0a3dbffcc4d5ae"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	newPassword := "secret1"
	// import 导入
	account, err := ks.Import(jsonBytes, password, newPassword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createKs()
	importKs()
}
