package main

import "fmt"

func main() {
	var a [5]int
	var b [2]string
	c := [2]string{"tuesday","friday"}
	a[0]=1
	b[1]="monday"
	fmt.Println(a,b,c)
	// 长度 容量。数组的长度为容量，不足的地方用零值填充
	fmt.Println(cap(a))
	fmt.Println(len(a))
	fmt.Println(a[0:3])

	// for遍历
	for k,v := range a {
		fmt.Println(k,v)
	}

	// 多维数组, 多少个维度就写多少个[] 然后最后标记值的类型
	// 总共10个数组作为元素，每个数组里面有2个元素
	d := [10][2]int{
		{1,2},
		{1,2},
	}
	fmt.Println(d)

	// 总共5个数组作为元素，每个数组中有3个string类型的元素
	var e [5][3]string
	// 这种方式定义，需要如下进行赋值
	e[0][0] = "Monday"
	e[0][1] = "Monday1"
	e[0][2] = "Monday2"
	fmt.Println(e)


}

