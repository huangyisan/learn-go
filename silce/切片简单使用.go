package main

import "fmt"

// slice不可比较

// 一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。
// slice为引用类型
// 切片不需要定义长度

func main() {
	//[]里面不需要写长度
	a:=[]int{1,2,3}
	//使用make创建 var slice1 []type = make([]type, len,cap)
	var slice1 []int = make([]int, 10,20)

	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}