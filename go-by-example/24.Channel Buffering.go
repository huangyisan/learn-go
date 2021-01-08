package main

import (
	"fmt"
	"sync"
)

// 默认channel无缓冲，但可以在make的时候指定缓冲长度，这样在长度塞满之前就可以塞入数据了


func myPrint(str chan string, wg *sync.WaitGroup) {
	fmt.Println(<- str)
	// 每次执行完则wg.Done(), 则wg.Add个数自动-1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	bufC := make(chan string, 2)

	wg.Add(3)  // 有几个goroutine就写几

	go myPrint(bufC, &wg)
	go myPrint(bufC, &wg)
	go myPrint(bufC, &wg)



	bufC <- "a"
	bufC <- "b"
	bufC <- "b"

	// 等待goroutine全部执行完
	wg.Wait()


}