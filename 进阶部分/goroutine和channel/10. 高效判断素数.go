package main

import "fmt"

/*
要求统计1-8000的数字中，哪些是素数。
分配四个goroutine去完成。

思路 每个goroutine执行完 往exitChan写入，主线程不停读取exitChan
 */

// 向intChan放入1-8000个数字
func putNum(intChan chan int) {
	for i:=1; i<8000; i++ {
		intChan <- i
	}
	// 退出就关闭管道
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool

	// 使用for循环
	for {
		// 开始让flag为true, 期望为素数
		flag = true
		num, ok := <- intChan
		if !ok {
			// 取不到数据，则退出
			break
		}

		// 判断num是否是素数
		for i:=2; i<num ; i++ {
			// 如果能被整除，则不是素数
			if num % i == 0{
				flag = false
				break
			}
		}
		// 上面循环都没触发 if条件，则可以知道i为素数
		if flag {
			//将该数字放入到primeChan中
			primeChan <- num
		}
	}
	// 全部处理完
	fmt.Println("其中一个判断素数方法执行完毕。")

	// 通知exitChan，表示自己完成，但不能关闭primeChan，因为总共有四个在执行，其他可能还在往primeChan中写数据
	exitChan <- true

}


func main() {
	intChan := make(chan int, 1000)
	// 存放素数
	primeChan := make(chan int, 2000)

	// 退出标示管道
	exitChan := make(chan bool, 4)


	// 开启一个协程，向intChan放入 1-8000 个数字
	go putNum(intChan)

	// 开启四个协程，从intChan取出数据，并且判断是否是素数，如果是就放入到primeChan中
	for i:=0; i<4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 主线程进行处理，从exitChan中取出四次，则表示结束
	// 取出四次即可, 这里也可以使用协程
	go func() {
		for i:=0; i<4; i++ {
			<- exitChan
		}
		// 取出后，则对primeChan关闭
		close(primeChan)
	}()

	// 遍历primeChan得到结果
	for {
		res, ok := <- primeChan
		if !ok {
			break
		}
		fmt.Println("素数为", res)
	}

}