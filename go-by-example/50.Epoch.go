package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 秒
	timestamp:= now.Unix()
	// 纳秒
	nanos := now.UnixNano()
	fmt.Println(now)
	fmt.Println(timestamp)
	fmt.Println(nanos)

	// 毫秒 = 纳秒 / 1000000
	mil := nanos / 1000000
	fmt.Println(mil)

	// 将时间戳转换为可视化时间
	fmt.Println(time.Unix(timestamp, 0))
	// 如果为纳秒转换,则纳秒作为第二个参数
	fmt.Println(time.Unix(0, nanos))
}