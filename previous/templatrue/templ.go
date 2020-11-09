package main

import "fmt"

func main() {
	var beforeDay int = 97
	var weeks,day int
	weeks = beforeDay / 7
	day = beforeDay % 7
	fmt.Printf("还剩%v周，%v天放假", weeks, day)

}
