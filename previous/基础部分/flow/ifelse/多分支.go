package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入一个年龄")
	fmt.Scanln(&age)

	if age > 18 && age < 30 {
		fmt.Println("大于18岁")
	}else if age > 30 {
		fmt.Println("大于30岁")
	} else {
		fmt.Println("小于18岁")
	}
}