package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func main(){
	var str string
	var b2 bool = true
	var bl string
	c := 22269
	b := false
	fmt.Printf("文字为: %c\n", c)
	fmt.Printf("文字b的占用空间为: %v\n", unsafe.Sizeof(b))
	d := `aasdfasdf""asdfsadf`
	fmt.Println("这是字符串d", d)

	str = strconv.Itoa(c)
	fmt.Println(str)

	str = "true"
	bl = strconv.FormatBool(b2)
	fmt.Println(bl)
	fmt.Printf("类型为%T",bl)

	var str3 string = "123.456"
	var f1 float64

	f1,_ = strconv.ParseFloat(str3,32)
	fmt.Println(f1)
	var i = 99
	i++
	var ptr *int = &i
	fmt.Printf("ptr的内容为%v", *ptr)
	fmt.Printf("自增为%v", i)


}
