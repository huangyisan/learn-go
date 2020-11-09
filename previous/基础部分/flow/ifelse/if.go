package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入一个年龄")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("大于18岁")
	}else {
		fmt.Println("小于18岁")
	}

	if age1 := 20; age1 > 18 {
		fmt.Println("age1大于18")
	}
}