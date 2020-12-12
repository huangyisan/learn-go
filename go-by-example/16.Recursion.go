package main

import "fmt"

func recursion(a,b int) int {
	sum := a + b
	if sum < 10 {
		fmt.Println(sum)
		a ++
		b ++
		recursion(a,b)
	}
	return sum
}

func main() {
	recursion(1,1)
}