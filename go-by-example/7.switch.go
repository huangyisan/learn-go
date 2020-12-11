package main

import (
	"fmt"
)

func main() {
	a := 0
	// 经典switch
	switch a {
	case 1:
		fmt.Println("a=1")

	case 2:
		fmt.Println("a=2")
	case 0:
		fmt.Println("a=0")
		// 往下穿透
		fallthrough
	default:
		fmt.Println("a != 1 2 0")
	}

	// case中做条件判断, 等同if else
	switch  {
	case a == 0:
		fmt.Println("a == 0")
	case a == 1:
		fmt.Println("a == 1")
	default:
		fmt.Println("a != 1 0")
	}

	// switch进行断言 i.(type)
	whatamI := func(i interface{}) {
		switch t:= i.(type) {
		case int:
			fmt.Println("类型为int")
		case string:
			fmt.Println("类型为string")
		default:
			fmt.Println("未知类型",t)
		}

	}
	whatamI(23)
	whatamI("test")
}
