package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 创建一个目录
	err := os.Mkdir("subDir", 755)
	check(err)

	// RemoveAll 等于rm -rf 清空目录下所有内容
	defer os.RemoveAll("subDir")

	// 在subdir下创建文件，使用函数的方法
	createFile := func(name string) {
		str := []byte("")
		check(ioutil.WriteFile(name, str, 644))
	}

	createFile("subDir/myfile")

	// 使用MkdirAll函数实现类似mkdir -p 效果
	os.MkdirAll("subDir/A/B/C/D", 755)

	// ReadDir方法列出目录下的内容
	c, err := ioutil.ReadDir("subDir")
	check(err)
	// 对内容迭代，则可以获取到名称 大小等信息
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir())
	}

	// Chdir类似cd命令
	os.Chdir("subDir")
	c1, err := ioutil.ReadDir(".")
	check(err)
	for _, entry := range c1 {
		fmt.Println(entry.IsDir())

	}

	os.Chdir("../")
	// Walk 接受一个路径和回调函数，用于处理访问到的每个目录和文件。所以当遍历访问每个文件的时候需要做什么就可以写成函数。
	fmt.Println("get subdir size")

	err1 := filepath.Walk("subDir", getSize)
	check(err1)
}

func getSize(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(p, info.Size())
	return nil

}
