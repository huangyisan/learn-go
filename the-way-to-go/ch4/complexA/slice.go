package complexA

import "fmt"

func CapLen() {
	months := [...]string{1: "January", 2: "2", 3: "3", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(len(months), cap(months))
	fmt.Println(len(Q2), cap(Q2))
	fmt.Println(len(summer), cap(summer))

	// extend a slice (within capacity)
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
