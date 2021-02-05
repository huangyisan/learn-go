package main

import (
    "fmt"
    "os"
)

// 通过os.Args获取执行传入的参数

func main() {
    fmt.Printf("Execute file absolute dir is %s\n", os.Args[0])
    fmt.Printf("ALl the args are %s\n", os.Args[1:])
    fmt.Printf("Second args is %s\n", os.Args[2])
}