package main

import (
    "fmt"
    "os"
)

func main() {
    // 当执行os.Exit则程序退出， defer内容不会再被执行
    defer fmt.Println("exit!")
    // Exit里面表示退出状态码
    os.Exit(3)
}
