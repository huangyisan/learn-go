package main

import "fmt"

// string也是一个byte数组, 因此string也可以进行切片处理.
//

func main() {
	str := "hello@guigu"
//	使用切片获取到@guigu
	slice1 := str[5:]
	fmt.Println(slice1)
}