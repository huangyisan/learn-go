package main

import (
	"fmt"
	"time"
)

// select 开监听channel上的数据流动。
// select语法和switch相似，但使用存在很多限制，最关键的一点是"每个case语句里面必须是一个IO操作"
// 在一个select语句中，go会按顺序从头至尾评估每一个发送和接受的信号。
// 如果没有任意一条语句可以执行，则出现两种可能。1. 存在default，则执行default内容。2. 不存在default，则阻塞，直到至少有一个通讯可以进行下去。


// 例子: 使用select做超时需求。
// 原理：让select进行探测，其中一条语句为case，同时定时器做为channel读取，如果读取到了，则超时。

func main() {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num:=<-ch:
				fmt.Println("channel有数据", num)
			case <-time.After(3*time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()
	// 将quit丢去
	<- quit
	fmt.Println("程序结束")
}