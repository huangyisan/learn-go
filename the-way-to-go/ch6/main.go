package main

import (
	"fmt"
	"learn/Method"
)

func main() {
	p := Method.Point{1, 2}
	q := Method.Point{4, 6}
	fmt.Println(p.Distance(q))
	r := &Method.Point{1, 2}
	r.ScaleBy(1.0)
	fmt.Println(*r)
}
