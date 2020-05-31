package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	// 提前关闭intChan
	close(intChan)

	// 此时无法进行塞入数据

	// 遍历管道不能使用普通的for循环
	// 如果没有提前关闭管道，则最后会报错。
	for v := range intChan {
		fmt.Println(v)
	}
}