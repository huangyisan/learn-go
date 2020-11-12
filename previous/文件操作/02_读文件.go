package main

import (
	"os"
	"fmt"
)

func readFile(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件异常")
		return
	}

	defer f.Close()

	// 给定一个切片，将读取的内容放里面, 总大小为2k
	buf := make([]byte, 1024*2)
	// 读取到的内容会放入buf中， n为读取到的长度
	n, err := f.Read(buf)
	if err != nil {

		fmt.Println("读取失败")
	}
	// 使用string方法转换为字符串
	fmt.Println(string(buf[:n]))
}

func main() {
	path := "./writefile.text"
	readFile(path)
}