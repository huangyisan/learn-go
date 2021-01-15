package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 1. 使用WriteFile方法，直接把要写的内容写入到文件中
	str1 := []byte("hello world\n")
	err := ioutil.WriteFile("tmp/write", str1, 0644)
	check(err)


	// 2. 创建文件，然后往里面写入ascii byte切片
	f, err := os.Create("tmp/write_create")
	check(err)
	defer f.Close()
	// 注意，这里是ASCII字符 115 111 109 101 10 -> s o m e \n
	d2 := []byte{115, 111, 109, 101, 10}
	//Write写ascii
	n, err := f.Write(d2)
	check(err)
	fmt.Printf("总共写入%d个",n)

	// 可以使用WriteString写入string类型
	f.WriteString("hello world\n")

	// 将结果回写到磁盘上
	f.Sync()



	// 3. 使用带缓冲的Writer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("Hello go\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// flush到磁盘,buff对应的是用Flush
	w.Flush()


}