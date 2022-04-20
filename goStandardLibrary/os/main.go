package main

import (
	"fmt"
	"os"
)

func CreateFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("file name, %s", f.Name())
	}
}
