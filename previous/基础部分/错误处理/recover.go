package main

import "fmt"

func testa() {
	fmt.Println("testa")
}

func testb (x int) {
	// 在defer中设置recover
	// 如果出现异常，则打印异常
	defer func() {
		err :=recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var a [10] int
	a[x] = 11
}

func testc() {
	fmt.Println("testc")
}

func main() {
	testa()
	x:=20
	testb(x)
	testc()
}