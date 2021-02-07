package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

// go接受unix信号，类似SIGTERM SIGINT，然后进行处理

func main() {
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // 使用notify方法 监听信号， 如果监听到args标记的信号，则往sigs管道传入
    // Notify函数让signal包将输入信号转发到c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。
    // signal包不会为了向c发送信息而阻塞（就是说如果发送时c阻塞了，signal包会直接放弃）：调用者应该保证c有足够的缓存空间可以跟上期望的信号频率。对使用单一信号用于通知的通道，缓存为1就足够了。
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        // 取出sigs中的内容
        sig := <- sigs
        fmt.Println()

        fmt.Println(sig)
        // 写入done管道内
        done <- true
    }()
    fmt.Println("awaiting signal")
    // 当没接收到信号的时候，下面会被阻塞
    <-done
    fmt.Println("exiting")

    // 配合switch case的话就可以做出当监听到某个信号，则执行某个操作的效果了。

}