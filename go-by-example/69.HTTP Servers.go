package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request)  {
    fmt.Fprintf(w, "hello world")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _,v := range headers {
            fmt.Fprintf(w, "%s -> %s\n", name, v)
        }
        //fmt.Println(name)
        fmt.Println(headers)
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe("0.0.0.0:8889", nil)
}