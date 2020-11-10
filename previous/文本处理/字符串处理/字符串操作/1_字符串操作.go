package main

import (
	"fmt"
	"strings"
)

func main() {
	// contains 是否包含子串
	// 包含返回true，不包含返回false
	fmt.Println(strings.Contains("hello","llo"))

	//Joins组合 将多个字符串进行组合
	s := []string{"aa","xx","yy"}
	newS := strings.Join(s,"")
	fmt.Println(newS)

	//Index 查询某个字符串所在的位置, 如果找到，则返回位置索引，没找到则返回-1
	fmt.Println(strings.Index("hello","ll"))
	fmt.Println(strings.Index("hello","xx"))

	// repeat 重复某个字符串
	fmt.Println(strings.Repeat("re",3))

	//split 按照指定的分隔符进行分割, 返回一个slice
	fmt.Println(strings.Split("yes or no", " "))

	//Trim去掉两边的指定字符串
	fmt.Println(strings.Trim(":no or yes:",":"))

	//fields用空格进行分割所有字符串, 该功能类似为用split指定空格
	fmt.Println(strings.Fields("cc ss sdf ff"))
}