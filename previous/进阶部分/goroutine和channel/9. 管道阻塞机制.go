package main

import "fmt"

// 如果只往挂管道里面写数据二不读数据， 则管道满了后，就会出现阻塞而dead lock.阻塞代码位置为 intChan <- i
// 即便写入的速度大于读取的速度，也不会产生阻塞。这个解释器已经帮处理了。无需担心。

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
	intChan := make(chan int , 10)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	//go readData(intChan,exitChan)

	// 直到能有数据可以读出来为止
	for {
		_,ok := <- exitChan
		if !ok {
			break
		}
	}
}