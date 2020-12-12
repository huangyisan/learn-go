package main

import "fmt"

func main() {
	//数组不定义长度就是切片了
	// 切片用make来创建
	var a = make([]int, 10)
	a[0]=2

	// 或者直接赋值具体内容
	b := []int{1,2,3,4,5,655,7}
	fmt.Println(a)
	fmt.Println(b)

	// slice切片
	fmt.Println(b[2:5])
	// 获取长度和容量，切片的长度并不一定和容量相同
	d := b[:1]
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	// append返回一个新的slice
	e := append(a,10)
	fmt.Println(e)

	// copy 将一个切片赋值给另外一个切片
	var f = make([]int,len(a))
	copy(f,a)
	fmt.Println(f)

	// 循环切片
	for i:=0; i<len(a); i++ {
		fmt.Println(a[i])
	}

	for _,v :=range b {
		fmt.Println(v)
	}
}
