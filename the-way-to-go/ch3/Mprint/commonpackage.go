package Mprint

import (
	"bytes"
	"fmt"
	"strconv"
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

func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func StringToInt() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)
}

func IntToString() {
	//var err error
	var x int
	var y int64
	x, _ = strconv.Atoi("23")
	y, _ = strconv.ParseInt("123", 10, 64)
	fmt.Println(x)
	fmt.Println(y)
}
