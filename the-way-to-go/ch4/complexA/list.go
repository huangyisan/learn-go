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
