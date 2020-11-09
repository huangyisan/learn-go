package main

import "fmt"

func re() (a int, b int, c int) {
	a=1
	b=2
	c=3
	return
}

func main() {
	a,b,c := re()
	fmt.Println(a,b,c)
}