package main

import "fmt"


func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <-1
		}
		//关闭channel
		close(ch)

	}()

	for {
		// 可以用range遍历管道。只要管道内没数据了，则自动退出
		for v :=range ch {
			fmt.Println(v)
		}
	}
}