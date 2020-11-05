package main

import "fmt"

func main(){
	const (
		a = iota
		b = iota
	)
	var ch byte
	ch = 97
	fmt.Printf("ccc %c\n", ch)
	fmt.Printf("a = %d, b = %d",a,b)
}
