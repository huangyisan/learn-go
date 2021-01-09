package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// atomic提供的原子操作能确保任一时刻只要一个goroutine对变量进行操作，善用atomic能够避免程序中出现的大量锁操作

func main() {
    // 计数器无符号整型，永远是正整数
    var ops uint64

    var wg sync.WaitGroup

    for i:=0; i<50; i++ {
        wg.Add(1)
        go func() {
            for c:=0; c< 1000 ;c ++ {
                // 使用Add 自增，应用内存地址，自增1，
                // 该行为确保同一时间只有一个goroutine能够进行操作
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops", ops)

}