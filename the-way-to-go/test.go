package main

import "fmt"

//var a int = 5
type me int

var a int = 1
var b int = 2

const hardEight = (1 << 100) >> 97
const c = -1

//var b me = 8

func myprint() {
	//fmt.Println(b)
	fmt.Println(c)
	fmt.Print(a)
	fmt.Print(b)

	//fmt.Println(123123)
}
func main() {
	myprint()
}
