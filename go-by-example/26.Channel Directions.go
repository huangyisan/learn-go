package main

import "fmt"

// channel可以设置为只进不出,或者只出不进

// 只准进
func write(str string, strC chan<- string) {
	strC <- str
}

// 只准出
func read(strC <- chan string) string {
	out := <-strC
	return out
}

func main() {
	strC := make(chan string, 2)
	write("test string", strC)
	out := read(strC)
	fmt.Printf(out)
}