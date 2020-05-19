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
	fmt.Printf("年%v, 月%v, 日%v, 时%v, 分%v, 秒%v\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//	3. 格式化,和其他语言不一样,不是使用yyyy mm hh,而是使用固定的数值 2006/01/02 15:04:05, go语言中通过不同的数值来区分是属于什么字段.
	// 相当有意思 https://stackoverflow.com/questions/45160822/what-does-20060102150405-mean
	fmt.Println(now.Format("15:04:05 2006/01/02"))


}
