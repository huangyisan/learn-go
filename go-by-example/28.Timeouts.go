package main

import (
	"fmt"
	"time"
)

// select 也可以用作超时妙用
// time.After是一个返回chan,为time类型的函数

func main() {
	ch1 := make(chan string, 1)
	go func() {
		// 2s后往channel塞入数据
		time.Sleep(time.Second * 2)
		ch1 <- "one"

	}()

	select {
	case a :=<- ch1:
		fmt.Println(a)
		// 超时等待为1s 所以直接就超时了
	case <- time.After(time.Second):
		panic("执行超时")
	}
}