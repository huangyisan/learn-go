package main

import "fmt"
// map不可相互比较
// map在使用前一定要make
// map的key是不能重复的,如果重复,则以最后一个key-value为准
// map的value是可以相同的
// map的key-value是无需的
// make可以指定长度

// 1. 如果使用map来定义变量，则需要用make来初始化。

func main() {
	//使用map仅仅声明变量, 之后必须用make初始化，否则是个nil的map无法插入数据
	var myMap map[string]int

	// 使用make初始化变量，如果不初始化，则无法对该nil的map插入数据
	myMap = make(map[string]int)

	//插入kv
	myMap["name"] = 123
	fmt.Println(myMap)


	//	直接用make来初始化map
	yourMap := make(map[string]int)
	yourMap["name"] = 456
	fmt.Println(yourMap)

	//	如果用map声明变量并且定义变量内容，则可以插入值
	rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	rating["name"] = 2.3
	fmt.Println(rating)
}



