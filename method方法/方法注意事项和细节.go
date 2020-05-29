package main
// 结构体类型是值类型， 在方法调用中， 遵循值类型的传递机制， 是值拷贝传递方式。
// 如果希望在方法中修改结构体变量的值， 可以通过结构体指针的方式来处理。

// 对于自定义类型，首先确定该类型是否实现了 String() 方法，如果实现了，那么fmt.Println会默认调用String() 方法进行输出

import "fmt"

type Student struct {
	Name string
	Age int
}

// 实现String()后， fmt.Println就按照String()方法的内容进行输出了
func (s Student) String()  string {
	return fmt.Sprintf("名称:%s, 年龄:%d", s.Name, s.Age)
}

func main() {
	var s = Student{
		Name: "qcrao",
		Age: 18,
	}
	// 因为 Student 结构体没有实现 String() 方法，所以 fmt.Println 会利用反射挨个打印成员变量

	fmt.Println(s)
}


