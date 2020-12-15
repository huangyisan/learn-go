package main

import (
	"fmt"
	"time"
)

// select 如果没有default, 那么所有case对应起码有一个chan有数据才会继续执行, 但如果有了default,则select的case都等待,则执行default内容.

func main() {
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "hello world"
	}()

	select {
	case <- ch1:
		fmt.Println("已经读取ch1")
	default:
		fmt.Println("还没轮到读取chan就执行了default")
	}
}