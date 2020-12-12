
package main

import "fmt"

type Person struct {
	Name string
	Age int
	Addr string
}

func person(name,addr string, age int) *Person {
	return &Person{
		name,
		age,
		addr,
	}
}

func main() {
	a := Person{
		"tom",
		10,
		"EN",
	}

	fmt.Println(a.Name)

	// 使用指针的方式
	b := &Person{
		"jack",
		10,
		"EN",
	}
	fmt.Println(b.Name)


	p := person("nick","CN",19)
	c := &a

	fmt.Println(c.Name,p)



}

