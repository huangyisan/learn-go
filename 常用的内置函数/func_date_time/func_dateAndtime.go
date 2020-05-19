package main

import (
	"time"
	"fmt"
)

func main() {


	// 1. 获取当前的时间
	now := time.Now()
	fmt.Printf("now的值为%v, 类型为%T\n", now, now)

	//2. 获取年月日 时分秒
	fmt.Printf("年%v, 月%v, 日%v, 时%v, 分%v, 秒%v", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

}
