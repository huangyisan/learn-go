package main

import (
	"fmt"
	"runtime"
)
// 主函数如果执行完，此时协程还没执行完，则协程会退出。
// 使用Gosched函数可以让出时间片，让协程优先执行完。

func main() {
	go func() {
		for i := 1; i< 10 ; i++ {
			fmt.Println("go")
		}
	}()

	for i:= 1; i<= 2 ; i++ {
		// 让出时间片，让协程优先执行完。
		runtime.Gosched()
		fmt.Println("main go")
	}
}