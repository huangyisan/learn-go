package main

import (
	"fmt"
	"errors"
)

// 支持自定义错误，使用errors.New和panic内置函数。
// 1. errors.New("错误说明")， 会返回一个error类型的值，表示一个错误。
// 2. panic内置函数，接受一个interface{}类型的值（任何值）作为参数。可以接受error类型的变量，输出错误信息，并退出程序。

func readConf (name string) (err error) {
	if name == "cc" {
		return nil
	} else {
		// 如果不是cc 则返回一个错误类型，且内容为"错误"
		return errors.New("错误")
	}
}

func test02() {
	err := readConf("dd")
	if err != nil {
		//抛出错误，并且终止代码往后执行
		panic(err)
	}
	// 抛出错误了，不会再执行
	fmt.Println("错误后的代码1")
}

func main() {
	test02()
	fmt.Println("错误后的代码2")
}