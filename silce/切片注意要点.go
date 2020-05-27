package main

import "fmt"

// 切片定义后还不能直接使用, 因为本身是一个空的,需要让其引用一个数组,或者make一个空间供切片来使用

func main() {
	var slice3 []int
	slice3 = []int{100, 200, 300}

	slice3 = append(slice3, 100)

	fmt.Println(slice3)

//	切片追加切片
	slice3 = append(slice3, slice3...)
	fmt.Println(slice3)



}