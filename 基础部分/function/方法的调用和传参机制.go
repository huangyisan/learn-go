package main

import "fmt"

// 方法不单单会传入函数中的变量,也会把实例传进去.

// 因为struct为值类型, 所以将struct传入到方法中后, 会复制出一份.

type Person struct{
	Name string
}

func (persion Person) Say(str string) {
	fmt.Printf("%s says %s", persion.Name, str)
}

func main() {
	Lee := Person{
		Name:"Lee",
	}

	Lee.Say("hello everyone")
	slicea := []int{1,2,3,4}
	fmt.Println(len(slicea), slicea[4])
	stra := []string{"1","2","3","4"}
	fmt.Println(len(stra))
}