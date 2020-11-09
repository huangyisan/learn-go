package main

import (
	"os"
	"fmt"
)

func main() {
	//打开文件
	// os.Open 返回指针，和err
	file, err := os.Open("./Test1")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(file)

	//关闭文件
	c := file.Close()
	if c != nil {
		fmt.Println(c)
	}
}
