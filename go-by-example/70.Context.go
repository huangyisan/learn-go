package main

import (
    "fmt"
    "net/http"
    "time"
)

func hello (w http.ResponseWriter, req *http.Request) {
    ctx := req.Context()
    fmt.Println("start to handler a http request")
    defer fmt.Println("handler end")

    select {
    case <- time.After(time.Second * 5):
        fmt.Println("handler ...")
        fmt.Fprintf(w, "handler a request now...")
        // 客户端在等待数据的时候ctrl+c强制关闭连接
    case <-ctx.Done():
        err := ctx.Err()
        fmt.Println("server err:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe("0.0.0.0:8889", nil)
}