package Declarations

import "fmt"

func PrintPrinter() {
	x := 2
	p := &x

	*p = 8
	fmt.Println(*p)
	fmt.Println(x)
}
