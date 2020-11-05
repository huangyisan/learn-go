package main

import (
	"time"
	"fmt"
)

// goroutine中使用recover可以解决协程中出现panic导致程序崩溃的问题。
// 如果一个协程异常，不要影响其他协程和主线程。


func sayHello(){
	for i:=0; i<10; i++{
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}


func test() {
	// 使用defer + recover 捕获panic
	defer func() {
		if err:=recover(); err != nil {
			fmt.Printf("执行test()发生错误",err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"
}

func main()  {
	go sayHello()
	go test()

	for i:=0; i<10; i++{
		fmt.Println("main")
	}


}