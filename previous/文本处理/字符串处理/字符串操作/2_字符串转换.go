package main

import (
	"fmt"
	"strconv"
)

func main() {
	// AppendTP系列函数将整形布尔等转换为字符串后,添加到现有的字节数组中.
	// 因为是字节数组,所以如下定义
	s := make([]byte,0)
	s = strconv.AppendBool(s, true)
	// 第三个参数10表示以十进制的方式转换,120也就是120
	s = strconv.AppendInt(s, 120, 10)
	// 追加带引号的字符串内容
	s = strconv.AppendQuote(s, "abc")
	// 使用string()函数将切片转换为字符串
	fmt.Println("slice = ",string(s))


	// 其他类型转换为字符串
	var str string
	// 布尔类型转换成字符串
	str = strconv.FormatBool(true)
	fmt.Println(str)
	// float类型转换为字符串, 第一位为float的数值,第二位为表示格式,第三位为精度,第四位为位数(float64)
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(str)
	// 整形转换  第一位为整型数字,第二位为进制
	str = strconv.FormatInt(999,2)
	fmt.Println(str)
	// 常用Itoa进行整型转字符串
	str = strconv.Itoa(100)
	fmt.Println(str)


	// 字符串转其他类型
	var flag bool
	// 该方法会返回两个值
	flag, _ = strconv.ParseBool("true")
	fmt.Println(flag)

	//


}
