package main

import "fmt"

/*
管道声明
var 变量名 chan 数据类型
var Num chan int  (管道存放int)
var Map chan map[string]string (管道存放map[string]string)
var person chan *Person
var person2 chan Person
 */

func main() {
	// 创建三个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan的值为%v\n", intChan)

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 4. 管道的长度在make的时候就被决定了。管道满了继续写，则报错deadlock
	fmt.Printf("长度为%v, 容量为%v\n", len(intChan), cap(intChan))

	// 5. 从管道读取数据

	var num2 int
	// 丢弃一个数据
	<- intChan
	num2 = <- intChan
	fmt.Println(num2)
	fmt.Printf("长度为%v, 容量为%v", len(intChan), cap(intChan))

	// 6.在没有协程情况下，如果管道数据全部取完， 继续取，则报错deadlock

}