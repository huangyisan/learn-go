package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	count :=0
	rand.Seed(time.Now().Unix())

	for {
		n := rand.Intn(100) + 1
		fmt.Println(n)
		count +=1
		if n == 99 {
			break
		}
	}
	fmt.Println(count)
}
