package main

import (
	"fmt"
	"time"
)

// 对channel初始化指定大于0个容量， 为有缓冲channel。
// 有缓存channel可以放入小于等于容量的个数，放满如果没有被消费，继续放入则阻塞。




func main() {
	// 指定空间为3， 所以该ch为有缓冲channel.
	ch := make(chan int, 3)

	// 打印出来，可以看到 空间为3
	fmt.Printf("len(ch) = %d\n", cap(ch))

	go func() {
		for i:= 0; i<3; i++ {
			fmt.Printf("子协程 %d \n", i)
			ch <- i
		}
	}()

	time.Sleep(time.Second)

	for i:=0; i<3; i++ {
		num := <-ch
		fmt.Printf("num = %d\n", num)
	}
}