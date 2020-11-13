package main

import "fmt"


// 定义读channel
func readOnly(ch <- chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

// 定义写channel
func writeOnly(ch chan <- int) {
	for i:=0; i<5; i++ {
		ch <- i
	}
	// 写完后关闭channel
	close(ch)
}


func main() {
	// 双向channel
	var ch = make(chan int, 5)

	go writeOnly(ch)
	go readOnly(ch)
}