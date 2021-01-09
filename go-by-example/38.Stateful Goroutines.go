package main
// 使用内建协程和通道的同步特性来达到互斥锁的效果
//  Go 共享内存的思想是，通过通信使每个数据仅被单个协程所拥有，即通过通信实现共享内存。 基于通道的方法与该思想完全一致！
import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key int
    resp chan int
}

// 定义writeOp结构体
type writeOp struct {
    key int
    val int
    resp chan bool
}

func main() {
    var readOps uint64
    var writeOps uint64

    // 定义reads chan， 里面存放readOp类型
    reads := make(chan readOp)
    writes := make(chan writeOp)

    go func() {
        var state = make(map[int]int)

        // 无限select，等待reads或者writes chan内数据
        for {
            select {
            // 读的每一个数据都是readOp结构体类型
            case read := <- reads:
                read.resp <- state[read.key]
            case write := <- writes:
                state[write.key] = state[write.val]
                write.resp <- true
            }
        }
    }()

    // 读操作
    for r :=0 ; r <100; r++ {
        go func() {
            for  {
                read := readOp{
                    key: rand.Intn(5),
                    resp: make(chan int),
                }
                reads <- read
                // 读出resp
                <- read.resp
                atomic.AddUint64(&readOps, 1)
            }

        }()
    }

    // 写操作
    for w:=0; w<10; w++ {
        go func() {
            for {
                write := writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(10),
                    resp: make(chan bool),
                }
                writes <- write
                // 结果resp
                <- write.resp

                atomic.AddUint64(&writeOps, 1)

            }

        }()
    }
    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOpsFianl", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOpsFinal", writeOpsFinal)

}