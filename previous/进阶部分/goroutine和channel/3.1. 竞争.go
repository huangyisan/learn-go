package main

import (
	"fmt"
	"time"
)

// 当多个协程调用一个资源的时候就会发生竞争.
// printHello和printWorld对printByte方法产生了竞争
func printByte(s string) {
	for _, v := range s {
		fmt.Printf("%c", v)
		time.Sleep(time.Second)

	}
}

func printHello() {
	s := "hello"
	printByte(s)

}

func printWorld() {
	s := "world"
	printByte(s)
}

func main() {
	//打印的时候产生了竞争. 使得printHello还没执行完毕, 资源就被printWorld抢走,
	go printHello()
	go printWorld()


	for {

	}

}