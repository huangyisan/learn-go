package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {
    str := "helloooo world"
    p("Contains", s.Contains(str, "el"))
    p("Count", s.Count(str, "l"))
    p("HasPrefix", s.HasPrefix(str, "wo"))
    p("HasSuffix", s.HasSuffix(str, "lo"))
    p("Index", s.Index(str, "l"))
    p("Join", s.Join([]string{"a","b"}, "|"))
    p("Repeat", s.Repeat(str, 2))
    // 替换一位
    p("Replace", s.Replace(str, "o", "0", 1))
    // -1 表示全部替换
    p("Replace", s.Replace(str, "o", "0", -1))
    p("Split", s.Split(str, " "))
    p("ToLower", s.ToLower(str))
    p("ToUpper", s.ToUpper(str))
    p()

    p("Len: ", len(str))
    // string转ASCII
    p("Char:", str[1])
}