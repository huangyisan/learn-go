package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 如果没有提前随机种子，则每次得到的数值都一样。
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    // 小数随机 rand.Float64 返回一个64位浮点数 f，且 0.0 <= f < 1.0
    // 5.0 <= f < 10.0
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    // 默认情况下，种子相同，产生的结果相同，如果要产生不同的结果，则需要先随机生成种子
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    // 使用新生成的随机种子进行获取随机数
    fmt.Println(r1.Intn(100))
    fmt.Println(r1.Float64())
}