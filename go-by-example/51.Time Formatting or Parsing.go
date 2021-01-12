package main

import (
	"fmt"
	"time"
)

// 2006-01-02 15:04:05

func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// 使用 2006-01-02 15:04:05 的格式来format时间
	p(t.Format("15:04"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	// 自定义时间格式来解析
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)


	// 可以使用对每个时间字段进行填充的方式来获取时间
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
