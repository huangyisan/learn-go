package main

import (
	"fmt"
	"time"
)

func plus() {
	for i:=10; i>0; i-- {
		time.Sleep(time.Second)
		fmt.Printf("goroutine %d\n",i)
	}
}

func main() {
	// 使用go 匿名函数内写执行函数
	go func() {
		plus()
	}()
	// 或者go 直接跟执行函数
	go plus()
	for i:=10; i>0; i-- {
		time.Sleep(time.Second * 2)
		fmt.Println(i)
	}


}