package Declarations

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func Incr() {
	v := 1
	incr(&v)
	fmt.Println(v)
}
