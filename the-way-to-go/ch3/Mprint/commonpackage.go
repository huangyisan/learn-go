package Mprint

import (
	"fmt"
	"strings"
)

func CommonPackage() {
	fmt.Println(basename("a/b/c/ccc.go"))
}

//func basename(s string) string {
//	for i := len(s) - 1; i >= 0; i-- {
//		fmt.Printf("%c\n", s[i])
//		if s[i] == '/' {
//			s = s[i+1:]
//			break
//		}
//	}
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] == '.' {
//			s = s[:i]
//			break
//		}
//	}
//	return s
//}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	fmt.Println(slash)
	s = s[slash+1:]
	//if dot := strings.LastIndex(s, "."); dot >= 0 {
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func StringToSlice() {
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(b)
	fmt.Println(s2)
}

func StringJoin() {
	str := []string{"Geeks", "For", "Geeks"}
	fmt.Println(strings.Join(str, "-"))
}
