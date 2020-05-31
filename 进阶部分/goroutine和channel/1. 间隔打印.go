package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i:=1; i<10; i++ {
		fmt.Println("hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test() //开启一个协程 关键字 go
	for i:=1; i<10; i++ {
		fmt.Println("hello man " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}