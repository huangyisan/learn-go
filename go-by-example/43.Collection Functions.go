package main

import (
    "fmt"
    "strings"
)

// 内联操作，不创建或调用帮助函数

// Index 返回目标字符串 t 在 vs 中第一次出现位置的索引， 或者在没有匹配值时返回 -1
func Index(vs []string, t string) int {
    for i,v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Any 如果切片 vs 中的任意一个字符串满足函数f，则返回 true。
func Any(vs []string, f func(string) bool ) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func main() {
    strSlice := []string{"a","b","c"}
    fmt.Println(Index(strSlice, "a"))
    fmt.Println(Any(strSlice, func(s string) bool {
        return strings.HasPrefix(s, "a")
    }))


}