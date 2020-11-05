package main

import (
	"fmt"
	"os"
)

// os.Args是一个string的切片， 用来存储所有命令行的参数

func main() {
	fmt.Println("总参数为", os.Args)
	// 脚本名称为第一个元素，也就是[0]
	fmt.Println("脚本名称为", os.Args[0])
	// 脚本执行跟随的参数从第二个元素开始，也就是从[1]开始
	fmt.Println("第一个参数为", os.Args[1])
	fmt.Println("第二个参数为", os.Args[2])
	fmt.Println("第三个参数为", os.Args[3])
}
