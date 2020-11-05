package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
//	读取文件并显示在终端,带缓冲区的方式
	file, err := os.Open("./进阶部分/文件操作/Test")
	if err != nil {
		fmt.Println(err)
	}
	//	关闭文件, 否则会有内存泄露
	defer file.Close()

//	创建一个 *Reader, 默认缓冲区为4096
	reader := bufio.NewReader(file)

	// 循环读取文件内容
	for {
		// 读取的内容存放在str, 读取错误存放在err中
		str, err := reader.ReadString('\n') //读到换行符就结束,表示按照行读取.
		fmt.Print(str)  //不是用Println是因为读的时候会把\n读进去, Println表示输出后换行.
		if err == io.EOF { // io.EOF表示读取到文件末尾
			break //则跳出读取
		}
	}
	fmt.Println("文件读取结束")



}