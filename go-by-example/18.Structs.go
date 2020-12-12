package main

import "fmt"

type Person struct {
	Name string
	Age int
	Addr string
}

func main() {
	a := Person{
		"tom",
		10,
		"EN",
	}

	fmt.Println(a.Name)

	b := &Person{
		"jack",
		10,
		"EN",
	}
	fmt.Println(b.Name)
}