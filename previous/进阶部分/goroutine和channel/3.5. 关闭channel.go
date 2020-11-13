package main

import "fmt"

// 关闭channel
// close(channelName)
// 如果对一个已经关闭了的channel写入数据，则会抛出错误

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
		// <-ch 操作返回两个值，一个为channel里面的内容，一个为状态。
		// 如果ok为true，则表示channel未关闭， 反之则表示channel关闭。
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else{
			break
		}
	}
}