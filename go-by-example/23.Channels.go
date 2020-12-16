package main

import "fmt"

func main() {
	chanStr := make(chan int)

	go func() {
		chanStr <- 1
	}()

	a := <- chanStr
	fmt.Println(a)

}