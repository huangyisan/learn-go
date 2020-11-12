package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccc")
	// 终止此函数所在的协程，且以下内容将不再执行，ddddd和eeeee不会被打印。
	runtime.Goexit()
	fmt.Println("dddddd")
	defer fmt.Println("eeee")

}

func main(){

	go func() {
		fmt.Println("aaa")
		test()
		// test函数中的Goexit导致该协程被终止，所以bbbb也不会被打印出来
		fmt.Println("bbbbb")
	}()

	for {

	}
}
