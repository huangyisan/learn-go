package main

import "fmt"

// 在函数中，需要创建资源，文件句柄，数据库连接等，在执行完毕后，关闭这些资源。go提供了defer(延时机制)


func sum(n1 int, n2 int) int {
	// 当执行defer时，会将defer后面的语句压入到单独的栈(defer栈)
	// 当函数执行完毕后，在从defer栈中，按照先入后出的方式，执行
	defer fmt.Println("第一个数为", n1) // 最后第一执行
	defer fmt.Println("第二个数为", n2) // 最后第二执行
	res := n1+n2
	fmt.Println("结果为",res)
	return res
}

func main() {
	sum(1,2)
}