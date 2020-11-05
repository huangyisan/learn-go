package main

import (
	"io/ioutil"
	"fmt"
)

// 一次性读取整个文件适合读取不大的文件.方法为 ioutil.ReadFile
// 成功的调用返回的err为nil而非EOF, 因为本函数定义为读取整个文件.

func main() {
	file := "/Users/huangyisan/Desktop/github/hys/learn-go/进阶部分/文件操作/Test"
	content, err := ioutil.ReadFile(file)
	if err !=nil {
		fmt.Println(err)
	}
	// 将读取到的内容显示到终端
	// 将byte转换为string
	fmt.Printf("%v", string(content))
	// 因为没有显式的open,所以也不需要显式的close关闭文件,这两个操作已经封装到ReadFile里面了

}