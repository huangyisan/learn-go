package main

import "fmt"

func main() {
    // panic后不会执行接下来的内容
    panic("a problem")

    fmt.Println("after probelm")
}