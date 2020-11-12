package main

import (
	"fmt"
	"time"
)
// channel也是一种数据类型,需要先定义, 全局定义,让函数使用.

var ch = make(chan int)

func printByte(s string) {
	for _, v := range s {
		fmt.Printf("%c", v)
		time.Sleep(time.Second)


	}
}

func printHello() {
	s := "hello"
	printByte(s)
	// 当执行完后, 往管道放入666
	ch <- 666

}

func printWorld() {
	// 如果ch中没有数据,则阻塞,获取到数据后则执行下面代码
	<- ch
	s := "world"
	printByte(s)

}

func main() {
	//当hello打印完后,往channel塞入数据, 然后printWorld的阻塞被释放,继续打印world, 解决了竞争.
	go printHello()
	go printWorld()


	for {

	}

}