package main

import "fmt"

func main() {
	// if
	age := 18
	if age < 8 {
		fmt.Println("xiao")
	}
	if age1 := 20; age1 < 30 {
		fmt.Println("da")
	} else {
		fmt.Println("xioao")
	}

	type Gender int8

	const (
		MALE Gender = 1
		FEMALE Gender = 2
	)
	// switch
	garder := MALE

	switch garder {
	case MALE:
		fmt.Println("MALE")
		fallthrough
	case FEMALE:
		fmt.Println("FEMALE")
		fallthrough
	default:
		fmt.Println("unknow")
	}
	// case做表达式， switch不需要写变量
	switch {
	case MALE > 1:
		fmt.Println("da")
	}

	// for
	num := 10
	for i:=1; i< num; i++ {
		fmt.Println(num)
	}

	s := []int{1,2,3,4,5}
	for _, v := range s {
		fmt.Println(v)
	}

	

}

