package main

import (
	"fmt"
	"time"
)

func woker(id int, job chan int, result chan int)  {
	for e := range job {
	fmt.Printf("worker: %d, process %d\n",id,e)
	time.Sleep(time.Second)
	result <- e
	}
}

func main() {
	job := make(chan int,100)
	result := make(chan int,100)

	// 开启3个并发， 再job没有数据的时候会阻塞
	for w:=0; w<3; w++ {
		go woker(w,job,result)
	}

	// 给job写入工作数量
	for i:=0; i<9;i++{
		job <- i
	}
	// 写完后关闭job channel
	close(job)

	// 读取result，由内容就读取，直到读完为止
	for i:=0; i<9; i++ {
		<-result
	}
}