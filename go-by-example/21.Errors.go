package main

import (
	"errors"
	"fmt"
)

// go抛出错误的两个方法
// 1. errors.New()
// 2. fmt.Errorf()

func raiseError() error {
	return errors.New("一个致命的错误")
}


func fError() error {
	return fmt.Errorf("另一个致命的错误")
}



// 只要实现了Error() string函数，则实现了error接口。
// error 代码
/*
type error interface {
	Error() string
}
 */

type errorInfo struct {
	str string
}

func (s errorInfo) Error() string {
	return fmt.Sprintf("这是一个自定义的错误，内容为: %s", s.str)
}

func main() {
	if e := raiseError(); e != nil {
		fmt.Println(e)
	}

	if e2 := fError(); e2 != nil {
		fmt.Println(e2)
	}

	err := &errorInfo{
		"数据为空",
	}

	fmt.Printf("%s",err)
}