package main

import "fmt"

// channel可以被关闭,如果关闭了,则无法继续对channel进行写入.

func main() {
	ch1 := make(chan string, 5)

	ch1 <- "one"
	ch1 <- "two"
	ch1 <- "three"

	//关闭管道
	close(ch1)

	a, more := <- ch1
	a, more = <- ch1
	a, more = <- ch1
	// 关闭后, 即便一直读也不会报错
	a, more = <- ch1
	a, more = <- ch1
	a, more = <- ch1
	a, more = <- ch1

	// 当关闭管道,且管道内容全部被读取后,more的值为false.反之只要两条件有一条不符合都为true
	if more {
		fmt.Println(a)
		fmt.Println(more)
	}else {
		fmt.Println("读取完毕")
		fmt.Println(more)
	}

	// 对关闭的管道继续写入数据,则报错
	ch1 <- "four"
}