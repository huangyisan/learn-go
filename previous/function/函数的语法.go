package main

import "fmt"

/*
func 函数名 (形参列表) （返回值列表） {
	执行语句
    return 返回值
}
*/

func cal(n1 int, n2 int, op byte) int {
	var res int
	switch op {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	}
	return res
}

func main() {
	fmt.Println(cal(1,2,'+'))
}