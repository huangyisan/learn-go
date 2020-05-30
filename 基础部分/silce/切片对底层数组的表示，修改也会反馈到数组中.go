package main

import (
"fmt"
)

func main() {
	//darr是数组
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// 将darr[2:5] 定义赋值给dslice作为切片
	dslice := darr[2:5]
	fmt.Println("array before",darr)
	// 对[2:5]切片的修改，会反应到数组上
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after",darr)
}
