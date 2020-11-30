package main

import "fmt"
// err :=recover()  recover()能捕获到产生的异常, 所以只要err != nil, 则表示出现异常.
// recover()常常和defer func() {}配合使用
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