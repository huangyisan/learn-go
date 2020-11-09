package main

import "fmt"

// 实际开发中，可能没我们想象中那么容易确定何时关闭管道，此时可以使用select

func main()  {
	intChan := make(chan int,10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d",i)
	}

	// 实际开发中，可能没我们想象中那么容易确定何时关闭管道，此时可以使用select。
	for {
		select {
			case v:= <- intChan : // 这里 如果intChan一直没关闭， 也不会阻塞。 会到下一个case中去取。
				fmt.Println(v)
			case v:= <- stringChan :
				fmt.Println(v)
			default:
				fmt.Println("全部取完了")
				// 结束退出， 或者break 配合label
				return
		}
	}

}
