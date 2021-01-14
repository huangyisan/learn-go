package main

import (
    "fmt"
    "strconv"
)

//使用stconv包可以对字符串进行数字转换处理

func main() {
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // 打印出的都是十进制数， 第二位base，表示用何种进制去解析，比如下面为16进制的123，其对应的十进制为291
    // 如果base为0则表示自动推断类型
    i, _ := strconv.ParseInt("123", 16, 64)
    fmt.Println(i)

    // ParseInt 可以自动识别带转换数为何种类型
    d, _ := strconv.ParseInt("0off", 0 , 64)
    fmt.Println(d)

    // 识别成无符号整型
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // Atoi直接转换成十进制
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // 如果传入的内容无法被转换，则报错
    _, err := strconv.Atoi("ss")
    fmt.Println(err)

}
