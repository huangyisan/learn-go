package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

type Rope string

func main() {
	for i := 0; i < 10; i++ {
		abc := rand.Int()
		fmt.Println(abc)
	}
	for i := 0; i < 8; i++ {
		a := rand.Intn(1100)
		fmt.Println(a)
	}
	fmt.Println()
	times := int64(time.Now().Nanosecond())
	rand.Seed(times)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(8))
	}
	var ch int = '\u0041'
	print(unicode.IsDigit(ch))
}
