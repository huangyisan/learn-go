package main

import "fmt"

func multivalue() (int, int) {
	return 7,8
}

func main() {
	a,b := multivalue()
	fmt.Println(a,b)
}