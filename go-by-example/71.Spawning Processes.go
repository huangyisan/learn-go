package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

// go调用外部进程，比如调用shell进程

func main() {
    // exec.Command 可以帮助我们创建一个对象，来表示这个外部进程。
    currentDir := exec.Command("ls")
    // 使用Output()获取输出内容
    info, err := currentDir.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("current dir")
    fmt.Println(string(info))


    // 搜集外部输入的stdin，然后对输出stdout搜集结果
    grepCmd := exec.Command("grep", "hello")

    // 创建grepIn stdin pipeline
    grepIn, _ := grepCmd.StdinPipe()
    // 创建grepOut stdin pipeline
    grepOut, _:= grepCmd.StdoutPipe()

    // 命令执行开始
    grepCmd.Start()
    // 往stdin pipeline写入
    grepIn.Write([]byte("hello world\nhellos"))
    // 关闭stdin pipeline
    grepIn.Close()
    // 从grepOut读取内容赋值给grepByte
    grepByte, _ := ioutil.ReadAll(grepOut)

    // 等待命令执行结束
    grepCmd.Wait()

    fmt.Println(string(grepByte))

    // exec.Command()执行命令必须传入多个参数，如果只想传入一个命令的字符串，则使用bash -c
    longString := exec.Command("bash","-c","ls -rlth")
    longStringOut, _ := longString.Output()
    fmt.Println(string(longStringOut))
}
