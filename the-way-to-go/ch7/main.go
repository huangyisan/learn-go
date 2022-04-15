package main

import (
	"fmt"
	"learn/InterFace"
)

var c InterFace.ByteCounter

func main() {
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
