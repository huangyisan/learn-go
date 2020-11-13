package main

import (
	"fmt"
	"time"
)

// 对channel初始化不指定空间， 或者空间为0的 为无缓存channel。
// 无缓存channel放入一个元素后将无法继续放入，如果继续放入，则会阻塞。
// 只有将channel中的元素取出后，才可以继续放入。




func main() {
	// 指定空间为0， 所以该ch为无缓存channel.
	ch := make(chan int, 0)

	fmt.Printf("len(ch) = %d\n", cap(ch))

	go func() {
		for i:= 0; i<3; i++ {
			fmt.Printf("子协程 %d \n", i)
			// 因为是无缓存channel，被阻塞，无法继续放入，只有等下面取出后才能继续放入。
			ch <- i
		}
	}()

	time.Sleep(time.Second)

	for i:=0; i<3; i++ {
		num := <-ch
		fmt.Printf("num = %d\n", num)
	}
}