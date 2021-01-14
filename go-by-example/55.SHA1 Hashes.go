package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {
    // 给定字符串，转换的时候需要转成[]byte类型
    str := "hello world"

    // 先生成一个散列对象
    s := sha1.New()

    // 对需要sha1的字符串转换后调用Write方法
    s.Write([]byte(str))

    // 求sum
    // Sum 得到最终的散列值的字符切片。Sum 接收一个参数， 可以用来给现有的字符切片追加额外的字节切片：但是一般都不需要这样做。
    bs := s.Sum(nil)

    // 用16进制字符串打印 %x
    fmt.Printf("%x\n", bs)
}