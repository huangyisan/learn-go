package Mprint

import "fmt"

func StringPrint() {
	s := "hello world"
	fmt.Println(s)
	fmt.Printf("%c,%c", s[0], s[2])
	fmt.Println()
}
