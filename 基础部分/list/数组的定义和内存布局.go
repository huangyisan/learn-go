package main

import "fmt"

// 数组定义
// var 数组名 [长度]数组类型

func main()  {
	var intArry [3]int
	//intArry[0] = 1
	//intArry[1] = 2
	//intArry[2] = 3
	// 数组的内存地址指向该数组收个元素的内存地址
	fmt.Printf("intArry地址为%v\n", intArry)
	fmt.Printf("数组首地址为%p\n", &intArry[0])
	// 数组第二个元素地址为第一个元素地址加上该数组类型占用的字节数,比如当前为int类型,则占用8个字节,那么第二个元素地址为第一个元素地址+8.后面的元素依次增加.
	fmt.Printf("第二个元素地址为%p\n", &intArry[1])
	fmt.Printf("第三个元素地址为%p\n", &intArry[2])
}

