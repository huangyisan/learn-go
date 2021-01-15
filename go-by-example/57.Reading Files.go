package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "strings"
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

    /*
        1. 将文件整个读入内存
        2. 按字节数读取   给定一个slice填充字节数，然后再给定一个slice用来append每次读到的数据， err == io.EOF则break
        3. 按行读取      通过ReadString来实现遇到'\n'作为分隔符来读文件，err == io.EOF则break
    */
    // 1. 全部读入方法1
    file, err := os.Open("tmp/dat")
    check(err)
    defer file.Close()

    // ReadAll方法读取
    content, err := ioutil.ReadAll(file)
    // content为[]byte类型， 使用string()方法处理
    fmt.Println(string(content))

    // 全部读入方法2 则为本篇开头使用ReadFile方法实现。


    // 2. 按字节读文件
    file2, err := os.Open("tmp/dat")
    check(err)
    defer file2.Close()

    // 使用默认buff大小，返回一个NewReader
    r := bufio.NewReader(file2)

    // 开设一个[]byte，用于最后存放数据
    chunks := make([]byte,0)

    // 开始一个[]byte, 用来临时存放,数据，临时个1024byte个数据，也就是每次读取1024byte
    buf := make([]byte, 1024)

    // for 无限循环进行读取
    for {
        // 该Read()读到最后返回的n为0，然后抛出一个io.EOF报错, 将读到的数据存入buf中，每次存入buf指定的大小
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }
        // 将buf内元素一个个打散append进chunks中
        chunks = append(chunks, buf[:n]...)

    }
    fmt.Println(string(chunks))


    // 3. 按行读取， 一行结束的标记为\n
    file3, err := os.Open("tmp/dat")
    check(err)
    defer file3.Close()

    buf2 := bufio.NewReader(file3)
    for {
        // ReadString 读取指定分隔符， 然后返回内容以及分隔符，这边指定\n为分隔符
        line, err := buf2.ReadString('\n')
        // 因为ReadString会返回内容以及分隔符'\n'，所以每次输出都会多一个空行，用TrimSpace删除空行
        line = strings.TrimSpace(line)
        //fmt.Println(strings.TrimSpace(line))
        fmt.Println(line)

        // 如果有报错，则判断是否为io.EOF，表示读完
        if err != nil {
            if err == io.EOF {
                break
            } else {
                fmt.Println("读取错误")
                panic(err)
            }
        }
    }
}