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

	//结构体也可以通过new来申请, 开辟一块新地址.这种方式的book2是指针类型
	book2 :=new(books)
	book2.page=22
	book2.name="yisan2"
	printBook(book2)


}

// 传入结构体的指针，
func printBook(book *books) {
	fmt.Println(book.page)
	fmt.Println(book.name)
}
