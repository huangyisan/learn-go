package main

import "fmt"

/*
开启一个writedata协程，向管道写入50个整数。
开启一个readdata协程，向管道读取writedata的数据
主线程等待writedata和readdata工作完后再退出
 */

//  如果告诉主线程已经读结束了？ 设计一根新的管道，当readData读取完后，向新管道写入一个值并关闭。主线程for循环读取该管道，当读到写入的值，则结束。



func writeData(intChan chan int) {
	for i :=1 ; i<=50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println(i)
	}

	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Println("读到数据", v)
	}
	// 读取任务完成，写入exitChan内true
	exitChan <- true
	close(exitChan)
}




func main() {
	// 创建两个管道
	intChan := make(chan int , 50)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)

	// 直到能有数据可以读出来为止
	for {
		_,ok := <- exitChan
		if !ok {
			break
		}
	}
}