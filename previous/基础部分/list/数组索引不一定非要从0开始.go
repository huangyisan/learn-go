package main

import "fmt"

func main() {
	// 直接从99索引值给定值为1，,70索引为2, 其他索引(0-70, 71-98)的为零值。
	// 99和70为数组索引的值.不要和python的字典搞混.
	r:=[...]int{99:1, 70:2}
	fmt.Println(r[99])
	fmt.Println(r[70])
	//索引0为零值，int的零值为0
	fmt.Println(r[0])
}
