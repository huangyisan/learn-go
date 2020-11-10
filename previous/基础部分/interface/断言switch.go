 package main

import "fmt"

// interface的变量可以存储任意实现了该interface的数值。空接口可以存储任何东西。
// 通过断言，可以反向知道，该变量里面保存的对象是哪个类型。

// 1. 使用comma-ok 断言

//Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
//如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。

// 如果一个变量类型为空接口， 那么里面可以包含任何类型， 如果不用断言，则无法区分

type blank interface {

}

type List []blank

type Person struct {
	Name string
	Age int
}

// 给person设定一个Pprint方法
func (self Person) Pprint() {
	fmt.Printf("名称为: %s\n年龄为: %d\n", self.Name, self.Age)
}

func main() {
	// 在list中混入了多种类型， 需要用断言才能将类型甄别开。
	list := make(List, 3)
	list[2] = 1
	list[1] = "name"
	list[0] = Person{Name:"yisan", Age:29}

	for _, element := range list{
		fmt.Printf("值为%v, 类型为%T\n",element ,element)
		// 断言写法 使用switch
		switch value:= element.(type) {
		case int:
			fmt.Printf("类型为%T, 值为%d\n", value, value)
		case string:
			fmt.Printf("类型为%T, 值为%s\n", valuze, value)
		case Person:
			persion := Person{Name:"huang", Age:23}
			persion.Pprint()
		}
	}
}
