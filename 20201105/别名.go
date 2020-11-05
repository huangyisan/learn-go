package main

import "fmt"

func main() {
	type bigint int64
	var a bigint
	fmt.Printf("%T", a)

}
