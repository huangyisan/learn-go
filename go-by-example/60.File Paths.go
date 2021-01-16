package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

// filepath 包为 文件路径 ，提供了方便的跨操作系统的解析和构建函数； 比如：Linux 下的 dir/file 和 Windows 下的 dir\file 。

func main() {
    // Join可以传入多个参数，并且进行路径的拼接。使用Join，而非手工写死"/"或者"\"
    pathSlice := []string{"dirA", "dirB", "dirC", "dirD", "filename"}
    p := filepath.Join(pathSlice...)
    fmt.Println("p:", p)

    // Join会删除多余的/
    fmt.Println(filepath.Join("abcDir//","/def"))

    // Dir和Base可以分别获取目录路基和文件名
    fmt.Println(filepath.Dir(p))
    fmt.Println(filepath.Base(p))

    // Split返回dir和base
    fmt.Println(filepath.Split(p))

    // IsAbs判断是否为绝对路径, 会和当前系统进行真实比对得到是否为绝路径
    fmt.Println(filepath.IsAbs("/abc/file"))
    fmt.Println(filepath.IsAbs("cc/abc"))

    // 使用Ext分割扩展名,返回扩展名
    filename := "readme.md"
    ext := filepath.Ext(filename)
    fmt.Println(ext)

    // 使用TrimSuffix获得除开扩展名的文件名
    fmt.Println(strings.TrimSuffix(filename,ext))

    // Rel 查询两者路径的相对路径， 如果查询不到则返回错误
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
    
    // 查询不到 返回错误
    rel, err = filepath.Rel("a/b", "/c/a/b/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

}
