package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main() {
	// 定义一个存放任意数据的类型管道，三个数据
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"

	tom := Cat{
		Name:"tom",
		Age:8,
	}

	allChan <- tom
	// 丢弃前两个
	<- allChan
	<- allChan
	jack := <- allChan

	fmt.Printf("类型为%T\n",jack) //打印出结构体类型


	// fmt.Printf("名字为%v",jack.Name) //此时执行出现报错 (type interface {} is interface with no methods)

	//解决方法是先进行类型断言
	a := jack.(Cat)
	fmt.Printf("名称为%v", a.Name)
}