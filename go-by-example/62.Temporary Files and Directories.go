package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // 使用ioutil.TempFile进行创建文件， dir为路径位置,如果为空则根据不同系统创建在临时文件所应该在的目录，比如win则在C:\Users\username\AppData\Local\Temp\下
    // pattern表示文件名以指定字符串开头的，且生成的文件名保证不会重复
    f, err := ioutil.TempFile("", "sample")
    check(err)
    // file name: C:\Users\xxx\AppData\Local\Temp\sample122907643
    fmt.Println("file name:", f.Name())

    // defer保证最后清理生成的临时文件
    defer os.Remove(f.Name())

    // 创建临时目录使用ioutil.TempDir
    dir, err := ioutil.TempDir("","mytmpdir")
    check(err)
    fmt.Println("tmp dir name:", dir)

    // defer保证最后清理目录
    defer os.RemoveAll(dir)

    // 将得到的临时目录和需要生成的文件拼接后写入内容
    fname := filepath.Join(dir, "myfile")
    err = ioutil.WriteFile(fname, []byte{1,2}, 0666)
    check(err)


}