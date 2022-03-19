package Mprint

import "fmt"

func ComplexPrint() {
	x := 1.2i
	fmt.Println(x)

	var y complex128 = complex(3, 4)
	fmt.Println(y)
}
