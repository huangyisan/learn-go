package main

import (
	"fmt"
	"time"
)

// 若主函数执行比goroutine快， 此时会结束代码执行。
// 使用channel可以等待goroutine执行结束完毕。


func do(done chan bool) {
	for i := 0; i<3; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	// 执行完，则往管道写入true
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("我是主函数")
	go func() {
		do(done)
	}()

	// 读取到done内容后才结束，当无法读取的时候会一直阻塞
	<- done
}