package complexA

import (
	"crypto/sha256"
	"fmt"
)

func Mylist() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(len(a) - 1)

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int
	q = [3]int{1, 2, 3}
	fmt.Println(q[2])

	r := [...]int{1, 2, 3, 4}
	fmt.Println(r[len(r)-1])

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	// 这里直接给的索引,而不是map
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
}

func ListCompare() {
	a := [2]int{2, 3}
	b := [...]int{2, 3}
	c := [...]int{4, 5}
	fmt.Println(a == b, b == c, a == c)
}

func Sha256() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))

	fmt.Printf("%x\n%x\n%t\n%T %T", c1, c2, c1 == c2, c1, c2)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func ChangeZero() {
	ptr := [32]byte{1: 'c'}
	fmt.Println(ptr)
	zero(&ptr)
	fmt.Println(ptr)
}

func SliceAppend() {
	var runes []rune
	for _, r := range "hello,世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func NoemptyF() {
	data := []string{"one", "", "", "three", "four"}
	//fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		fmt.Println(s)
		if s != "" {
			strings[i] = s
			i++
		}
	}
	fmt.Println(strings)
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func SamulateStack() {
	var stack []int
	stack = make([]int, 0, 0)
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	pop1 := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(pop1)
	fmt.Println(stack)
	pop2 := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(pop2)
	fmt.Println(stack)
}

func SliceRemove() {
	fmt.Println(sliceRemove([]int{1, 2, 3, 4, 5}, 2))
}

// [1 2 3 4 5]
func sliceRemove(slice []int, i int) []int {
	fmt.Println(slice[i:], slice[i+1:])
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice)-1]
}

func CopySlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 6, 8, 9, 9}
	copy(s1, s2)
	fmt.Println(s1)
}
