package main

import (
    "fmt"
    "path/filepath"
)

// filepath 包为 文件路径 ，提供了方便的跨操作系统的解析和构建函数； 比如：Linux 下的 dir/file 和 Windows 下的 dir\file 。

func main() {
    // Join可以传入多个参数，并且进行路径的拼接。使用Join，而非手工写死"/"或者"\"
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)
}
