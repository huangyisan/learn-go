package main

import "fmt"

var ch = make(chan string, 10)

func download(url string) {
	ch <- url
	fmt.Printf("download %s", url)
}

func main() {
	for i:= 0; i<10; i++ {
		go download("/api"+string(i+1))
	}

	for i :=0; i<10; i++ {
		msg := <- ch
		fmt.Println("finish download %s", msg)
	}
	fmt.Println("done")
}
