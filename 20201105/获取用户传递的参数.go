package main

import (
	"os"
	"fmt"
)

func main() {
	//接受用户传递的参数，都是以字符串的方式传递的
	list := os.Args
	n := len(list)
	fmt.Println(n)

	for _,v := range list{
		fmt.Println(v)
	}
}


