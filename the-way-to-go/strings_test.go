package main

import (
	"fmt"
	"strings"
)

func hasPrefix(mystring, prefix string) bool {
	if strings.HasPrefix(mystring, prefix) {
		return true
	}
	return false
}

func main() {
	dString := "bacdef"
	jPrefix := "bac"
	res := hasPrefix(dString, jPrefix)
	fmt.Println(res)
}
