package main

import (
	"reflect"
	"fmt"
)

// 通过反射修改num int的值


func reflectInt(b interface{}) {
	rVal := reflect.ValueOf(b)
	// 查看kind, 此时是指针
	fmt.Println(rVal.Kind())

	// rVal提供了setInt()来修改原来的值, 此时的rVal是ptr指针类型,不能直接使用.
	// 需要先使用Elem()方法来获取指针指向的值, 然后再进行setInt()
	rVal.Elem().SetInt(20)
	//rVal.SetInt(20)
}

func main() {
	var num int
	num = 10

	// 因为要修改 所以要传递地址. int默认是值拷贝
	reflectInt(&num)

	fmt.Println("修改后的num为",num)

}