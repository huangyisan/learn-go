package Mprint

import "fmt"

func MyPrint() {
	a := 0007
	fmt.Printf("%d, %[1]o, %[1]X, %#[1]x\n", a)
}

func ChPrint() {
	b := '国'
	fmt.Printf("%c, %[1]d, %[1]x\n", b)
}
