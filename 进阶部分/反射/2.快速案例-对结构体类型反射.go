package main

import (
	"reflect"
	"fmt"
)

// 结构体, interface{}, reflect.Value三者相互转换的基本操作

type student struct {
	Name string
	Age int
}


// 反射
func reflectTest02(b interface{}) {
	// 通过反射, 获取传入变量的type, kind, 值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	// 虽然输出为int, 但并不是int类型, 使用%T就能看到,该类型为reflect.Type
	fmt.Println("rTyp =",rTyp)
	fmt.Printf("rTyp 类型为%T\n",rTyp)


	// 2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	// 获取kind
	fmt.Println("kind = ", rVal.Kind())
	fmt.Printf("kind类型为 = %T\n", rVal.Kind())

	// 虽然输出值为100, 但并不是100, 使用%T就能看到,该类型为reflect.Value
	fmt.Println("rVal = ", rVal)
	fmt.Printf("rVal 类型为%T\n",rVal)

	// 转成原本的值
	//realVal := rVal.
	//fmt.Printf("类型为%T\n", realVal)

	// 转成interface{}
	iV := rVal.Interface()
	fmt.Println("Interface{} = ",iV)

	// 将interface{} 通过断言转成int
	num2 := iV.(student)
	fmt.Println("name为", num2.Name)


	fmt.Printf("类型为%T", num2.Name)



}


func main() {
	// 定义一个student实例
	stu := student{
		Name:"yisan",
		Age:20,
	}


	reflectTest02(stu)
}