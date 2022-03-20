package Mprint

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Friday
	Saturday
)

func IotaPrint() {
	fmt.Println(Saturday)
}
