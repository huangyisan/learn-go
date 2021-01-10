package main

import (
    "fmt"
    "os"
)

type point struct {
    x int
    y int
}

func main() {
    p := point{4, 5}

    fmt.Printf("%v\n", p)

    // %+v可以打印更加详细的内容，结构体则会打印出字段名
    fmt.Printf("%+v\n", p)
    // %#v 打印出源码片段
    fmt.Printf("%#v\n", p)
    // %T 打印值类型
    fmt.Printf("%T\n", p)
    // %t 传入true或false，转string类型
    fmt.Printf("%t\n", true)
    // 使用Sprintf则返回格式化的结果，类似python的format
    a := fmt.Sprintf("%t", true)
    fmt.Printf("%T", a)
    // 格式化整形，"%d" 十进制格式
    fmt.Printf("%d\n", 100)
    // 格式化整形，"%b" 二进制格式
    fmt.Printf("%b\n", 100)
    // 格式化整形，"%c" 转ASCII
    fmt.Printf("%c\n", 100)
    // 格式化整形，"%x" 十六进制格式, 小写字母a-f
    fmt.Printf("%x\n", 100)
    // 格式化整形，"%x" 十六进制格式, 大写字母A-F
    fmt.Printf("%X\n", 100)
    // 格式化整形，"%f" 浮点进制格式
    fmt.Printf("%f\n", 20.1)
    // 格式化整形，"%e" 科学计数法，e小写, 传入float类型
    fmt.Printf("%e\n", 100.1)
    // 格式化整形，"%E" 科学计数法，E大写， 传入float类型
    fmt.Printf("%E\n", 100.2)
    // %s 字符串输出
    fmt.Printf("%s\n", "hello world")
    // %q 带双引号的字符串输出
    fmt.Printf("%q\n", "hello world")
    // %x 当对应字符串的时候，输出base16编码，每个字节使用2个字符表示
    fmt.Printf("%x\n", "hex this")
    // %p输出指针的值
    fmt.Printf("%p\n", &p)

    // 格式调整
    // %后面使用数字，表示宽度。默认情况右对齐并且用空格填充
    fmt.Printf("|%6f|%10f|\n", 1.1,20.2)
    // %宽度.精度。 用来格式化结果的宽度和精度
    fmt.Printf("|%6.3f|%10.4f|\n", 1.1,20.2)
    // 在宽度左边加 "-" 表示左对齐
    fmt.Printf("|%-6.3f|%-10.4f|\n", 1.1,20.2)

    // 使用Fprintf可以格式化输出, 指定io.Writers类型
    fmt.Fprintf(os.Stdout, "123")
}