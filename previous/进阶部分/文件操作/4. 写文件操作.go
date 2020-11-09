package main

import (
	"os"
	"fmt"
	"bufio"
)

// os.OpenFile
// OpenFile 方法第二个参数为打开方式, 第三个参数为文件属性(仅在unix系统下有效)
/*
具体打开方式有
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)

文件属性为unix 下  0600 0655 等表示方法
 */

func main() {

	filePath := "./进阶部分/文件操作/Test-write"
	file , err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 写入内容
	str := "hello world\n"

	// 创建一个带缓存的writer, 将file作为写入对象
	writer := bufio.NewWriter(file)

	for i:=0; i<5; i++ {
		writer.WriteString(str)
	}

	// 写完后, 从buffer  flush到磁盘中
	writer.Flush()

}
