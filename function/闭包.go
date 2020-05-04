package main

import "fmt"

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		return n+x
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}