package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pid=%d, time=%d", os.Getpid(), time.Now().Unix())
	})
	http.ListenAndServe(":9999", nil)
}
