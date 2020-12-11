package main

import (
	"fmt"
	"math"
)

func main() {
	const a int = 1
	const b string = "today"

	fmt.Println(a,b)

	const c float64 = float64(a)


	fmt.Println(math.Sin(c))
}
