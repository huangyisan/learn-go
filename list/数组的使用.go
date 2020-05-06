package main

import "fmt"

func main() {
	var a [3]int = [3]int{1,2,3}
	for i,v :=range a{
		fmt.Println(i,v)
	}
	for _,c := range a{
		fmt.Println(c)
	}
//	 使用...设定不定长数组，具体长度根据赋值的元素个数决定
	b := [...]int{1,2,3,4,5}
	fmt.Println(b[1:3])
	fmt.Println(b[1])



}