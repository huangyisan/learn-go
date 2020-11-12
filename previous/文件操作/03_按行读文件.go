package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func readFileLine(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件异常")
		return
	}

	defer f.Close()
	//创建一个缓冲区，内容放里面

	r := bufio.NewReader(f)
	// 无限循环
	for {
		//遇到\n则结束读取, \n也会被读取
		buf, err := r.ReadBytes('\n')
		if err != nil {
			// 当err内容为io.EOF 则文件结束 停止读取
			if err == io.EOF {
				break
			}
			fmt.Println("读取错误")
		}
		fmt.Println("err = ", string(buf))
	}

}

func main() {
	path := "./writefile.text"
	readFileLine(path)
}