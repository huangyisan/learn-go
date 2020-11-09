package main

import "fmt"

func add(a int) (err error) {
	if a == 3{
		fmt.Println("我是if")
		return
	}
	defer fmt.Println("我是if return后面的defer")
	return
}

func main() {
	add(3)

}