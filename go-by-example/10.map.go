package main

import "fmt"

func main() {
	// make后才能使用
	var a = make(map[string]int)
	a["s"] = 1
	fmt.Println(a)
	
	//直接赋值
	b := map[int]string{1:"aa", 2:"bb"}
	fmt.Println(b)

	// 获取值
	b1 := b[1]
	fmt.Println(b1)

	// 删除一个值
	delete(b,2)
	fmt.Println(b)
}
