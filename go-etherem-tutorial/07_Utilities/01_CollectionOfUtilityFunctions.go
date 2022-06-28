package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	// 验证是否有效以太坊地址util.IsValidAddress
	valid := hexutil.IsValidAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	fmt.Println(valid) // true

}
