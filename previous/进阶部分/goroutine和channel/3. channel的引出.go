package main

import (
	"fmt"
	"time"
)

// 使用goroutine实现1-200的阶乘，并且放入map中，最后打印显示。

// 1. 编写函数，计算各个数的阶乘，并放入map中。
// 2. 启动协程多个， 统计的结果放入map。
// 3. map应该是全局，这样才能让协程都来使用。

var (
	myMap = make(map[int]int, 10)
)


// 计算n的阶乘，将结果放入myMap中
func test(n int) {
	res := 1
	for i :=1; i<=n; i++ {
		res *= i
	}

	// 将res结果放入myMap
	myMap[n] = res
}

// 存在的问题
/*
写入map存在竞争，不安全。
可能主线程比协程提前结束，从而导致协程被迫停止。
 */
func main() {
	//启动多个协程
	for i := 1; i<=200; i++ {
		go test(i)
	}

	// 防止主进程先结束，手工10s等待。
	time.Sleep(time.Second * 10)

	for k,v := range myMap {
		fmt.Printf("key is %v, value is %v\n",k,v)
	}

}