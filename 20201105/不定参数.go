package main

import "fmt"

func MyFunc(args... int) {
	fmt.Printf("%d and %d \n", args[0], args[1])
}

func MyFuncD(args... int) {
	for _, data := range args{
		fmt.Printf("data = %d\n", data)
	}
}

func MyFuncc(a int, args... int) {
	fmt.Println(a)
	fmt.Printf("%d and %d \n", args[0], args[1])
	//将全部参数传递给MyFuncD
	MyFuncD(args...)
	//args...其实是切片，所以可以将最后一个参数传递给MyFuncD函数
	MyFuncD(args[len(args)-1:]...)
}


func main() {
	MyFunc(1,2)
	MyFuncc(1,2,3)
}