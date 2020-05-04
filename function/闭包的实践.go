package main

import (
	"strings"
	"fmt"
)

//传入一个文件名，.jpg后缀，如果传入有，则返回文件名，如果没有，则加上后缀后返回。


func makeSuffix(suffix string) func(string) string {

	return func(name string) string{
		if strings.HasSuffix(name, suffix) {
			return name
		}else {
			name = name + suffix
			return name
		}
	}
}

func main() {
	var name string
	fmt.Println("请输入一个文件名")
	fmt.Scanln(&name)
	fmt.Println(makeSuffix(".jpg")(name))
}
