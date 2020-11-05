package main

import (
	"runtime"
	"fmt"
)

// runtime包存在函数获取当前机器逻辑cpu个数。
// 也可以设置程序最大使用多少个cpu进行运行。

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	// 自己设置使用多少个cpu
	// go1.8之后默认运行在多个cpu上，可以不用设置该参数。
	runtime.GOMAXPROCS(cpuNum - 1) //设置当前cpu逻辑数量-1
}
