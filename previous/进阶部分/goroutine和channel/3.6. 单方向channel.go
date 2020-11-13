package main

import "fmt"

// 在定义channel的时候，指定了方向，则该channel为单方向channel。
// 单方向channel分两种，一种是只读，一种是只写。
// <- chan 表示从管道出，为只读。 例子  var ch chan <- int
// chan <- 表示从管道入，为只写。 例子  var ch <- chan int
// 可以将双向channel转换为单向channel，但不可以将单向channel转换为双向channel

func main() {
	var n int
	// 创建一个channel，双向
	ch := make(chan int)

	// 双向隐式转换为单向写数据channel
	var writeCh chan <- int = ch

	// 只能读
	var readCh <- chan int = ch


	writeCh <- 666
	n = <- readCh
	fmt.Println(n)


}