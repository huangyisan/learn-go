package main

import "fmt"

func main() {
	countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}

	fmt.Println("原始 map")
	fmt.Println(countryCapitalMap)

	// 添加一个key value
	countryCapitalMap["China"] = "BeiJin"
	fmt.Println(countryCapitalMap)

	// 修改一个key value
	countryCapitalMap["Italy"] = "cc"
	fmt.Println(countryCapitalMap)

	//delete函数删除指定map的指定key, 如果key不存在,也不会报错.
	delete(countryCapitalMap,"France")

	fmt.Println("删除France后的 map")
	fmt.Println(countryCapitalMap)
//	如果要删除key,则遍历key,然后delete删除, 或者给原来的map make一个新的空间.

//	通过key访问，看是否存在该key，如果不存在，则返回value对应的零值
	value, ok := countryCapitalMap["Italy"]
	//存在italy， 返回Italy对应的value，以及true
	fmt.Println(value,ok)

	value,ok = countryCapitalMap["China"]
	//不存在，返回零值，string为空，ok为false
	fmt.Println(value,ok)


//	通过len查看map长度，也就是key的个数, 而不是cap
fmt.Println(len(countryCapitalMap))
}