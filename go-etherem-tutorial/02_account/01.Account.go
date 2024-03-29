package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// 要使用go-ethereum的账户, 则必须将地址转换为common.Address类型.
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
