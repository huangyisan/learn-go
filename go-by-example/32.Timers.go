package main

import (
	"fmt"
	"time"
)
// NewTimer只会执行一次
func main() {
	timer1 := time.NewTimer(time.Second * 2)

	// .C 表示使用Time类型的channel
	// 也就是2秒钟后出数据，等于等待2秒执行之后内容
	// 这种使用等同time.sleep
	<- timer1.C
	fmt.Println("2s passed")

	// 但是用NewTimer产生的定时器，可以在定时器失效前，用stop()取消他。
	timer2 := time.NewTimer(time.Second * 3)

	// 原本需要等待3s后执行的，被timer2.Stop()给停止了, 所以<- timer2.C就一直阻塞了，go协程下面就不会执行。
	go func() {
		<- timer2.C
		fmt.Println("3s passed")
	}()

	timer2.Stop()
	fmt.Println("stop timer")
}