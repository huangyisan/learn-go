package main

import "fmt"

func main() {
	// 直接从99索引值给定值为1，其他索引(0-98)的为零值。
	r:=[...]int{99:1}
	fmt.Println(r[99])
	//索引0为零值，int的零值为0
	fmt.Println(r[0])
}
