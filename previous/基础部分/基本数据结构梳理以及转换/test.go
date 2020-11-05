package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "123.33"

	num, _ := strconv.ParseFloat(str, 32)

	fmt.Printf("num值为%f, num类型为%T\n", num,num)
}
