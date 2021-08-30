package main

import "fmt"

func get(index int) (ret int) {
	defer func() {
		ret +=1
		fmt.Println("woshi difer")
		if r := recover(); r != nil {
			fmt.Println("some error occor")
			ret = -1
		}
	}()

	s := []int{1,2,3}
	return s[index]

}

func main() {
	fmt.Println(get(1))
	fmt.Println("finish")
}
