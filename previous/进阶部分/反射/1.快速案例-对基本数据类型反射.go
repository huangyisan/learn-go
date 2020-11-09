package main

import (
	"reflect"
	"fmt"
)

// 基本数据类型, interface{}, reflect.Value三者相互转换的基本操作


// 反射
func reflectTest(b interface{}) {
	// 通过反射, 获取传入变量的type, kind, 值
	// 1. 先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	// 虽然输出为int, 但并不是int类型, 使用%T就能看到,该类型为reflect.Type
	fmt.Println("rTyp =",rTyp)
	fmt.Printf("rTyp 类型为%T\n",rTyp)


	// 2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	// 虽然输出值为100, 但并不是100, 使用%T就能看到,该类型为reflect.Value
	fmt.Println("rVal = ", rVal)
	fmt.Printf("rVal 类型为%T\n",rVal)

	// 转成原本的值
	realVal := rVal.Int()
	fmt.Printf("类型为%T\n", realVal)

	// 转成interface{}
	iV := rVal.Interface()

	// 将interface{} 通过断言转成int
	num2 := iV.(int)

	fmt.Printf("类型为%T", num2)



}


func main() {
	// 定义一个int
	var num int = 100
	reflectTest(num)
}