package main

import "fmt"

func main() {
	str := "huangyisan"

	// 转换成byte数组, 可以处理英文和数组 ,[]byte按照字节来处理,而汉字占用3个字节
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

//	 转换成rune数组, 可以处理中文
	arr2 := []rune(str)
	arr2[0] = '黄'
	str = string(arr2)
	fmt.Println(str)

}