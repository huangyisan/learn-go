package main

import "fmt"

// 对数组进行指定索引以及对应的值进行初始化数组。

// 看着像map类型，但其实是数组类型

func main() {
	type Currency int

	const (
		// iota为0，下面的依次递增1， 所以作为了索引。其实是0,1,2,3索引值
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(symbol[USD])
}
