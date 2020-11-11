package main

import (
    "os"
    "fmt"
)

func writeString(path,str string) {
    // 创建文件
    f, err := os.Create(path)
    if err != nil {
        fmt.Println("创建文件出错")
        return
    }

    // 最后关闭文件
    defer f.Close()

    for i:=0; i<9; i++ {
        strCon := fmt.Sprintf("i = %d \n", i) + str
        // writeString写文件
        f.WriteString(strCon)
    }
}

func main() {
    path := "./writefile.text"
    str := "c"
    writeString(path, str)
}