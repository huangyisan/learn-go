package Mprint

import (
	"fmt"
	"unicode/utf8"
)

func PrintUnicodeLen() {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t %c\n", i, r)
		i += size
	}

	for i, r := range "hello world 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
