package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        var times int
        for {
            times ++
            fmt.Println(times)
            time.Sleep(time.Second * 2)
        }
    }()
    var output string
    fmt.Println(fmt.Scanln(&output))

}
