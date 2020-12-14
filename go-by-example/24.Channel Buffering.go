package main

import "fmt"

// 默认channel无缓冲，但可以在make的时候指定缓冲长度，这样在长度塞满之前就可以塞入数据了

func main() {
	bufC := make(chan string, 2)
	bufC <- "a"
	bufC <- "b"

	fmt.Println(<-bufC)
	fmt.Println(<-bufC)
}