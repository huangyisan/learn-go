package main

import "fmt"

// 不定参数
// ...跟类型连着写


func mulitArgs(i ...int) int  {
	fmt.Println(i)
	var sum int
	for _,v := range i {
		sum += v
	}
	return sum
}

func main() {
	res := mulitArgs(1,2,3)
	fmt.Println(res)

	// 传入的参数其实为slice类型
	s := []int{1,2,3,4,5}
	res1 := mulitArgs(s...)
	fmt.Println(res1)
}
