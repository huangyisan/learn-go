package main

import "fmt"

func add (str1,str2 *int) int {
	*str1 += *str1
	return *str1 + *str2
}

func main() {
	str1 := 1
	str2 := 1

	sum := add(&str1, &str2)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(sum)
}
