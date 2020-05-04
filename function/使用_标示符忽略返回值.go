package main

import "fmt"

func getsumc(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	// return不跟随任何内容，表示返回sum和sub
	return
}

func main() {
	res1,_ := getsumc(1,2)
	fmt.Println(res1)
}
