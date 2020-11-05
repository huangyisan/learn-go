package main

import "fmt"

// 每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被go运行框架调用，也就是说init会在main函数前调用


func init(){
	fmt.Print("我是init函数")
}


// main函数中不需要调用init函数。init函数在执行的时候会被自动调用。
func main() {
	fmt.Println("我是main函数")
}