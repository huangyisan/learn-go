package main

import "fmt"

// 使用range遍历channel
// 保证channel是关闭的情况下进行遍历

func main() {

	ch1 := make(chan string, 2)
	ch1 <- "string1"
	ch1 <- "string2"

	// 关闭后的channel才可以遍历，否则遍历完继续遍历就报错了
	close(ch1)

	// channel只有值，没有索引
	for v := range ch1 {
		fmt.Println(v)
	}

}

