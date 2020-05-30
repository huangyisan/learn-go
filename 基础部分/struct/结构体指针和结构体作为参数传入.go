package main

import "fmt"

// 可以通过结构体指针的方式访问

type books struct{
	name string
	page int

}

func main() {
	var book1 books

	book1.page=20
	book1.name = "yisan"

	//将地址作为参数传入
	//也是将结构体作为参数传入了函数中
	printBook(&book1)

}

// 传入结构体的指针，
func printBook(book *books) {
	fmt.Println(book.page)
	fmt.Println(book.name)
}
