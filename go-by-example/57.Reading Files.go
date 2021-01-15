package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // 读取全部内容
    dat, err := ioutil.ReadFile("tmp/dat")
    check(err)
    fmt.Println(string(dat))

    // 对文件打开的方式进行后续处理
    f, err := os.Open("tmp/dat")
    // 记得关闭文件
    defer f.Close()
    check(err)

    // 创建一个byte类型的slice, 读取5个byte, 只读5个byte
    b1 := make([]byte, 5)
    // 数据读入到b1中, 返回的n1为个数
    n1, err := f.Read(b1)
    check(err)
    // 通过string()方法将所有的byte slice转换后得到内容
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    // seek到文件指定位置,然后读取
    o2, err := f.Seek(1, 0)
    check(err)
    // 读取2个bytle
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: \n", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))



    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 3)
    // 最少从b3 byte slice读取3个
    n3, err := io.ReadAtLeast(f, b3, 3)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // 使用seek(0,0)可以将读游标回到起点重新开始读取.
    _, err = f.Seek(0, 0)
    check(err)
    r4 := bufio.NewReader(f)
    // Peek函数为指定读取多少内容, 这里是5个byte内容,类似实现了
    // b3 := make([]byte, 5)
    // n2, err := f.Read(b2)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))


    // go读取整个文件
    /*  1. 将文件整个读入内存
        2. 按字节数读取
        3. 按行读取

     */


}