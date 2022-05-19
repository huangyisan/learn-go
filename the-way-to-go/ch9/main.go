package main

import "fmt"

func IsPalindrome(s string) bool {
	strRune := []rune(s)
	for i := range strRune {
		fmt.Printf("%v\n", i)
		fmt.Printf("%c\n", strRune[i])
		fmt.Printf("%c\n", strRune[len(strRune)-1-i])
		if strRune[i] != strRune[len(strRune)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("是不是"))
}
