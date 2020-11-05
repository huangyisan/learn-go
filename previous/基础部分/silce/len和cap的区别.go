package main

import "fmt"

//切片的长度是切片中元素的数量。切片的容量是从创建切片的索引开始的底层数组中元素的数量。

func main() {
	//darr是数组
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// dslice为切片，引用了底层darr数组[2:5)的部分,索引开始为2，所以cap为len(darr)-2
	dslice := darr[2:5]

	fmt.Println(len(dslice)) //长度为3
	fmt.Println(cap(dslice)) //实际反馈的是底层数组的长度(容量)-索引开始位置


}
