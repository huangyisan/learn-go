package main

import "fmt"

var a = "G"

func main() {
	n() // G
	m() // O
	n() // O
	var cb int
	b := 0
	cb += 1
	fmt.Println(b)
	fmt.Println(cb)
	fmt.Println(1 << 10)
}

func n() {
	fmt.Println(a)
}

func m() {
	a = "O"
	fmt.Println(a)
}
