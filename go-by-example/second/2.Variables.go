package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	var b int = 2
	var c = 3
	var c1 byte = 'a'
	var d float32 = 3.2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(c1)
	fmt.Println(d)

	str1 := "golang"
	str2 := "Go语言"
	Typestr := fmt.Sprintf("%T", str1)
	fmt.Println(reflect.TypeOf(str1[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%s\n", Typestr)
	fmt.Println("str2 长度: ", len(str2))

	runeArr := []rune(str2)
	fmt.Println(str2[3])
	fmt.Println(runeArr[2])
	fmt.Println(string(runeArr[2]))
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())

	chlist := "好不好"
	for _,v := range chlist {
		fmt.Println(string(v))
	}

	var arr [5]int
	arr = [5]int{1,2,3,4,5}
	fmt.Println(arr)

	arr2 := [10]int{1,2}
	fmt.Println(arr2)
	for k,v := range arr2 {
		fmt.Println(k,v)
	}

	var slice1 []int
	slice1 = make([]int, 3, 5)
	fmt.Println(slice1)
	slice1 = append(slice1, 30,30,90)
	slice1[0] = 2
	fmt.Println(slice1, cap(slice1))
	comb := append(slice1, slice1...)
	fmt.Println(comb)

	var map1 map[int]string
	map1 = map[int]string{1:"cc", 2:"dd"}
	fmt.Println(map1)

	po := "golang"
	var p *string
	p = &po
	fmt.Println(*p)



}