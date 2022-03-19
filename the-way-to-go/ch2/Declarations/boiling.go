package Declarations

import "fmt"

const bolingF = 212.0

func PrintBioling() {
	var f = bolingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)
}
