package main

import "fmt"

func add(a int) (err error) {
	defer fmt.Println("我是defer")

	if a == 3{
		fmt.Println("我是if")
		return
	}
	return
}


func main() {
	add(3)

}