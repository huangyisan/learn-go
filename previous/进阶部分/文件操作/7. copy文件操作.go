package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

// 使用io包的Copy函数实现.

// 将Test文件copy到源目录,名称为Test-copy


func CopyFile(ori string, dst string) (written int64, err error) {
	// 打开源文件
	srcFile, err := os.Open(ori)
	defer srcFile.Close()
	if err !=nil {
		fmt.Println(err)
	}
	// 通过srcFile获取Reader
	newReader := bufio.NewReader(srcFile)


	// 打开目标文件, 因为目标文件可能不存在,所以用OpenFile函数
	file , err := os.OpenFile(dst, os.O_WRONLY | os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	newWriter := bufio.NewWriter(file)

	// 不全部读入内存，按缓冲区一定大小逐渐读取，防止内存溢出。
	count,err := io.Copy(newWriter, newReader)
	// Flush方法将数据从缓存刷入到文件中
	newWriter.Flush()
	return count, err
}

func main () {
	dst := "/Users/huangyisan/Desktop/github/hys/learn-go/进阶部分/文件操作/Test-copy"
	src := "/Users/huangyisan/Desktop/github/hys/learn-go/进阶部分/文件操作/Test"

	// 调用函数进行copy
	_, err := CopyFile(src, dst)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("copy成功")
	}
}
