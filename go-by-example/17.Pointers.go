package main

import "fmt"

// 使用指针可以对原值进行修改
func notPointer(i int)  {
	i = 0

}

func pointer(i *int)  {
	*i = 0

}

func main()  {
	// 通过函数修改i
	i := 4
	notPointer(i)
	fmt.Println(i)

	// 传入指针，所以i被修改了
	pointer(&i)
	fmt.Println(i)

}