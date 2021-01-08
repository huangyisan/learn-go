package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个缓冲为5的chan
	request := make(chan int,5)
	for i := 1; i <=5; i++ {
		request <- i
	}
	// 填装完毕后关闭request
	close(request)

	// 设定一个限制器，200ms, time.Tick()返回的是一个channel,每隔指定的时间会有数据从channel中出来
	limiter := time.Tick(time.Millisecond * 200)

	// 遍历request中的5个chan
	for req := range request {
		// 200ms产生一个，则读一个，然后继续执行，等于是sleep 200ms
		<-limiter
		fmt.Println("request", req, time.Now())
	}


	//给定一个burstyLimiter chan，用于突发三次限流
	burstyLimiter := make(chan time.Time, 3)

	// 填充3次数据
	for i :=0; i<3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// 每200ms往burstyLimiter填充数据，开始是无法填充的，因为burstyLimiter为满，需要等消费
		for t := range time.Tick(time.Millisecond * 2000) {
			burstyLimiter <- t
		}
	}()

	// 定义burstyRequest用来消费
	burstyRequest := make(chan int, 5)
	// 填充数据
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}

	// 模拟消费burstyRequest
	for req := range burstyRequest {
		//前面三次消费后，开始2000ms才往burstyLimiter填充数据，限流了
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
