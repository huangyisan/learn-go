package main

import "fmt"

func div(num1, num2 float32) (res float32) {
	res = num1 / num2
	//res2 = num1 % num2
	return res
}

func main() {
	fmt.Println(div(1,2))
}