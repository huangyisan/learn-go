package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var state = make(map[int]int)

    // 定义互斥锁
    var mutex = &sync.Mutex{}

    var readOps uint64
    var writeOps uint64

    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            // 一百个协程不停的执行以下内容
            for {
                key := rand.Intn(5)
                // 获得锁
                mutex.Lock()
                total += state[key]
                //fmt.Println("total:",total)
                // 释放锁
                mutex.Unlock()
                // 每次执行readOps计数cas方式+1
                atomic.AddUint64(&readOps, 1)

                time.Sleep(time.Microsecond)
            }
        }()
    }

    for w :=0; w <10 ; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(10)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Microsecond)
            }
        }()
    }

    time.Sleep(time.Second)
    // 载入操作能够保证原子的读变量的值，当读取的时候，任何其他 goroutine 都无法对该变量进行读写
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOpsFinal", readOpsFinal)
    wirteOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("wirteOpsFinal", wirteOpsFinal)

    mutex.Lock()
    fmt.Println("state", state)
    mutex.Unlock()

}