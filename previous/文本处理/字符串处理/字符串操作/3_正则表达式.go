package main

import (
	"regexp"
	"fmt"
)

func main() {
	testStr := "b2c,b2b bbc bbb"

	// 构造正则解析对象
	reg := regexp.MustCompile(`b\dc`)
	if reg == nil {
		fmt.Println("构造正则解析对象异常")
		return
	}

	// 进行匹配
	// -1 最大数量匹配, 如果是正数,则表示匹配个数
	res1 := reg.FindAllStringSubmatch(testStr, -1)
	fmt.Println(res1)
}