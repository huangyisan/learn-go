package main

import "fmt"

func test() {
	i := 1
	l := i
	x := []string{"a", "b", "c"}
	y := x
	fmt.Println(i)
	fmt.Println(l)

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%v\n", &i)
	fmt.Printf("%v\n", &l)

	fmt.Printf("%p\n", &x)
	fmt.Printf("%p\n", &y)

	i = 2
	x[1] = "vv"
	fmt.Println(l)
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	test()
}
