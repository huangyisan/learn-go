package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Printf("下载地址%s, 下载完成 \n", url)
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go download("/api")
	}
	wg.Wait()
	fmt.Printf("finish")

}