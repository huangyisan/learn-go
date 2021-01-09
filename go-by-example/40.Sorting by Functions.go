package main

import (
    "fmt"
    "sort"
)

// 自定义排序需要实现sort.Sort的interface的三个方法
// Len Less Swap
// https://studygolang.com/pkgdoc

/*
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
 */


type byLength []string

func (s byLength) Len() int {
    return len(s)
}

// 根据字符串长度来排序
func (s byLength) Less(i,j int) bool {
    return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i,j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {

    str := []string{"one","o","ne"}
    // 将str抓换成byLength类型
    strByLength := byLength(str)
    // 实现自定义排序
    sort.Sort(strByLength)
    //sort.Sort(byLength(str))
    fmt.Println(str)

}
