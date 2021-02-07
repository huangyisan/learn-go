package main

import (
    "flag"
    "fmt"
    "os"
)

// 使用flag包，实现命令行标志
// 基本的标记声明仅支持字符串、整数和布尔值选项。

func main() {
    // 字符串类型
    aString := flag.String("string", "abc", "input a string")
    // int类型
    aNumber := flag.Int("number", 1, "input a number")
    // bool类型
    aBool := flag.Bool("bool", true, "input bool")

    // 可以直接将var定义的对象进行参数化, 将其指针传入即可
    var name string
    flag.StringVar(&name, "name", "Jack","your name")

    // 所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
    flag.Parse()

    // 打印参数也是得用指针的方式
    fmt.Println(*aString)
    fmt.Println(*aNumber)
    fmt.Println(*aBool)
    fmt.Println(name)
    // 打印所有标志
    fmt.Println(flag.Args())

    fmt.Println(os.Args)
    fmt.Println(len(os.Args))
}

/*
$ ./65.Command-Line\ Flags.exe -h
Usage of xxxx\65.Command-Line Flags.exe:
  -bool
        input bool (default true)
  -name string
        your name (default "Jack")
  -number int
        input a number (default 1)
  -string string
        input a string (default "abc")

$ ./65.Command-Line\ Flags.exe -bool true -name abc -number 11 -string mystring
abc
1
true
Jack
[true -name abc -number 11 -string mystring]


 */