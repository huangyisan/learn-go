package main

import "fmt"

// switch语句默认添加了break语句
// case表达式可以多个，使用,间隔
// case表达式后面的值类型必须和表达式一致

// switch后面也可以不带表达式，此时作为if else来使用

// switch后面可以直接定义变量并且赋值，但不推荐。

func main() {
	var age int
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	switch age {
	case 19:
		fmt.Println("大于18")

	case 1:
		fmt.Println("小于18")
	//当age匹配5或者6或者7，则执行
	case 5,6,7:
		fmt.Println("5 6 7")
		//关键字fallthrough，表示穿透，执行完该case继续往下，无条件执行
		fallthrough
	case 99:
		fmt.Println("999")


	default:
		fmt.Println("等于18")

	}
}
