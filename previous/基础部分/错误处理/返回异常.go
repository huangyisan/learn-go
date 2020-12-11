package main

import (
	"errors"
	"fmt"
)

// 返回异常两种方式
// 1. fmt.Errorf
// 2. errors.New()

// 使用方法1
func div(left, right int) (c int, err error) {
	// 除数为0 直接返回错误
	if right == 0 {
		return 0, fmt.Errorf("除数不可以为0")
	}
	c = left / right
	return c, nil
}

// 使用方法2
func people(age int) (age1 int, err error) {
	if age == 0{
		return 0, errors.New("年龄不可以为0")
	}
	return age1, nil
}

func main() {
	c, err := div(1,0)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(c)
	}


	if age, err := people(0); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("年龄为%d", age)
	}
}