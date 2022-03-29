package main

import (
	"fmt"
	"learn/functions"
	"strings"
)

func add1(r rune) rune { return r + 1 }

func main() {
	//functions.CallchangeSlice()
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"

	// 匿名函数
	strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")

	functions.CallSquare()

}
