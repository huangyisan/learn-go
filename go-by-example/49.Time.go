package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 获取当前时间
	now := time.Now()
	p(now)

	// 提供日期时间时区等信息构建一个time
	then := time.Date(2020,01,01,23,23,12,500,time.FixedZone("CST",28800)) //offset为秒， 这边则定义时区为CST，东八区
	p(then)

	// 提取时间的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Date())
	p(then.Minute())
	p(then.Location())

	// 输出星期几
	p(then.Weekday())


	// 比较两天，某天在另外一天之前或者之后，或者两天为同一天
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub返回两个时间点的间隔时间
	diff := now.Sub(then)
	p(diff)

	// 获取间隔的具体时间段长度
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 使用Add 调整时间段 前移或者后移
	p(then.Add(diff))
	p(then.Add(-diff))

	// 打印unix时间
	p(time.Now().Unix())
}