package Declarations

import "fmt"

func New() {
	p := new(string)
	*p = "dog"
	fmt.Printf("this is %s", *p)
}
