package main

import "fmt"

// 空接口是一个万能接口，可以保存任意类型的值

func main() {
	var i interface{} = 1
	var k interface{} = "abc"
	fmt.Println(i,k)
}
