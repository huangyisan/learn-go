package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	file1 := "/Users/huangyisan/Desktop/github/hys/learn-go/进阶部分/文件操作/Test"
	file2 := "/Users/huangyisan/Desktop/github/hys/learn-go/进阶部分/文件操作/Test-write"

	data, err := ioutil.ReadFile(file1)
	if err != nil {
		return
	}

	err := ioutil.WriteFile(file2,data,0666)
	if err != nil {
		fmt.Println(err)
	}
}
