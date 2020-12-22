package main

import "fmt"

func main() {
	// 空结构体, map可以用来判断是否存在某个key,存在则返回true,不存在,则返回false
	aMap := make(map[string]struct{})
	aMap["a"] = struct{}{}
	aMap["b"] = struct{}{}
	aMap["c"] = struct{}{}

	_, ok := aMap["a"]
	fmt.Println(ok)
	_, ok = aMap["d"]
	fmt.Println(ok)

	// 空结构体相等
	a := struct {

	}{}
	b := struct {

	}{}

	fmt.Println(a == b)
	



}
