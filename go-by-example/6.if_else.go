package main

import "fmt"

func main() {
	// 单if
	a := 1
	if a == 1 {
		fmt.Println("a = 1")
	}

	// if - else
	if a == 0 {
		fmt.Println(" a = 0")
	}else {
		fmt.Println("a != 0")
	}

	// if - else if - else
	if a == 0 {
		fmt.Println("a = 0")
	}else if a == 1 {
		fmt.Println("a = 1")
	}else if a == 2 {
		fmt.Println("a = 2")
	}else {
		fmt.Println("a != 0 1 2")
	}
}