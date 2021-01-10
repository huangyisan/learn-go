package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    // 测试字符串是否符合表达式，返回bool和err
    match, _ := regexp.MatchString("^s.+d$","succeed")
    fmt.Println(match)

    // 获取一个正则对象，用于后面匹配等操作
    r, _ := regexp.Compile("p([a-z]+)ch")

    // 测试匹配，返回bool值
    fmt.Println(r.MatchString("peach"))
    // 返回匹配到的字符串
    fmt.Println(r.FindString("peach11 punch"))
    // 返回第一次匹配到的字符串的起始位索引和末尾位索引
    fmt.Println(r.FindStringIndex("peach11 punch"))
    // 返回完全匹配的字符串和组匹配的字符串，这里会返回 p([a-z]+)ch 和 ([a-z]+) 的信息。
    fmt.Println(r.FindStringSubmatch("peach11 punch"))
    // 返回完全匹配和局部匹配位置的起始位和结束位索引
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    // 返回指定匹配成功次数的结果，如果为-1，则全部返回匹配成功的结果
    fmt.Println(r.FindAllString("peach punch pinch", -1))
    // 返回指定匹配成功次数的结果的首末位索引，如果为-1，则全部返回匹配成功的结果首末位索引
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))
    // 使用[]byte配合match方法来实现MatchString一样的效果，返回为bool类型
    fmt.Println(r.Match([]byte("peach punch")))

    // 使用MustCompile, 如果异常则panic，使全局变量更加安全
    r = regexp.MustCompile("p([a-z]+)ch")
    // 使用子字符串替换匹配到的内容
    fmt.Println(r.ReplaceAllString("a peach","fruit"))

    // Func 变体允许您使用给定的函数来转换匹配的文本。

    in := []byte("a peach")
    // 匹配到的内容，传递给bytes.ToUpper函数进行处理
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))


}