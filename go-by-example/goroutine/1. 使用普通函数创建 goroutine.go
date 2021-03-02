package main

import (
    "fmt"
    "time"
)

func running() {
    var times int
    for {
        times ++
        fmt.Println(times)
        time.Sleep(time.Second * 2)
    }
}

func main() {
    go running()
    var output string
    fmt.Println(fmt.Scanln(&output))
}
