package main

import (
	"fmt"
	"time"
)

// select 可以让多个通道同时等待, 通道内有数据则立即执行

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var str string

	go func() {
		time.Sleep(time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	for i:=0; i<2 ;i++ {
		select {
		case str = <-ch2:
			fmt.Println(str)
		case str = <-ch1:
			fmt.Println(str)
		}
	}
}