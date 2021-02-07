package main

import (
    "fmt"
    "os"
    "strings"
)

// 使用go设置环境变量。
// os.Setenv设置变量，键值对方式
// os.Getenv获取键值对，没获取到则返回一个空字符串
// os.Environ列出所有环境变量的键值对, 键值对为slice，key value可以用=分割

func main() {
    os.Setenv("name","steam")
    name := os.Getenv("name")
    fmt.Println(name)

    notSetEnv := os.Getenv("notSetEnv")
    if notSetEnv == "" {
        fmt.Println("no set env")
    }

    allEnv := os.Environ()
    for _, v := range allEnv {
        fmt.Println(v)
        // 将内容以"="形式分割
        fmt.Println(strings.SplitN(v, "=", 2))
    }


}
