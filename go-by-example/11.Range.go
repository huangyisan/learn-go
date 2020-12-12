package main

import "fmt"

func main() {
	// 遍历切片
	a := []int{1,2,3,4,5}
	for _,v := range a {
		fmt.Println(v)
	}

	// 遍历字典
	b := map[string]string{"a":"one","b":"two"}
	for k,v := range b {
		fmt.Println(k, v)
	}

	// 迭代字符串, 输出assic字符
	c := "golang"
	for k,v :=range c {
		fmt.Println(k,v)
	}

	// 迭代中文，用range, 转换成[]rune， 在string转换回来
	d := "星期一"
	e := []rune(d)
	for _,v :=range e{
		fmt.Printf("%v",string(v))
	}
}