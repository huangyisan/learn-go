package main

import (
	"errors"
	"fmt"
	"os"
)

func hello(name string) error {
	if len(name) ==0 {
		return errors.New("a cuole ")
	}
	return nil
}

func main() {
	_, err := os.Open("cc")
	if err != nil {
		fmt.Println("error")
	}

	err = hello("")
	if err != nil {
		fmt.Println(err)
	}
}