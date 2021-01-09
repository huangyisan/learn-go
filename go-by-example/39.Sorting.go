package main

import (
    "fmt"
    "sort"
)

func main() {
    // 待排序字符串数组
    str := []string{"b","g","a","z","h"}
    // 原地排序，直接改变给定的切片，而非返回一个新值
    sort.Strings(str)
    fmt.Println("After sorting: ", str)

    // 待排序数字
    num := []int{4,2,33,37,1}
    // 原地排序
    sort.Ints(num)
    fmt.Println("After sorting: ", num)

    // 检查某个数组是否有序
    res := sort.IntsAreSorted(num)
    fmt.Println("is sorted?", res)
}