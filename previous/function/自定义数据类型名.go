package main

import "fmt"

// 对函数值进行命名， 这样的话 如果有多个返回值，返回值顺序无所谓(go可以返回多个数据)，只要名称对上即可。
func getsumc(n1 int, n2 int) (sum int) {
	sum = n1+n2
	return sum
}

// 函数为参数的时候。容易过长，可以type自定义类型。
type myFunctype func(int, int) int
//func myFuncc(funcvar func(int, int) int, n1 int, n2 int) int {
func myFuncc(funcvar myFunctype, n1 int, n2 int) int {
	return funcvar(n1,n2)
}


func main() {
	a := getsumc
	fmt.Printf("a的数据类型为%T, getsum数据类型为%T", a, getsumc)

	b1 := 1
	b2 := 2

	b := myFuncc(getsumc,b1,b2)

	fmt.Println(b)

//	给int取别名为myInt，数据类型虽然都为int，但对于go而言，myInt和int还是两个类型
//定义myInt为int类型
	type myInt int
	var num myInt
	num = 100
	// 以下写法会报错，虽然都是int类型，但myInt和int还是两个类型
	//var num1 int
	//num1 = num

	fmt.Println(num)
}
