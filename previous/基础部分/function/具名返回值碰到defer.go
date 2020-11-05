package main

import (
	"fmt"
)

func foo() (result string) {
	result = "ccc"
	defer func() {
		result = "Change World" // change value at the very last moment
	}()
	return result
}

func main() {
	fmt.Println(foo())
}
