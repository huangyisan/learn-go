package main

import "fmt"


func main() {
	// 1. 赋值； 判断条件； 变量改变
	for i:=10; i>0; i-- {
		fmt.Println("gogogo")
	}

	//2. 循环判断条件
	j := 0
	for j <= 10 {
		fmt.Println("ddddd")
		j++
	}

	//3. 无限循环
	k := 1
	for {
		if k < 10 {
			fmt.Println("ddddcc")

		}else {
			//跳出for循环
			break
		}
		k++
	}

	//	4. for range循环, 如下写法使用字符进行遍历，所以中文不会有问题。
	var str string = "hello 上海"
	for index, val:= range str {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}

	// 这种方式为通过字节遍历，因为中文占用utf8三个字节，所以遍历需要转换成rune切片。
	var str2 string = "hello 北京"
	str3 := []rune(str2)
	for i := 0; i<len(str3); i++ {
		fmt.Printf("%c \n", str3[i])
	}

}