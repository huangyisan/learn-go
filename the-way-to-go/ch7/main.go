package main

import (
	"bytes"
	"fmt"
	"io"
	"learn/InterFace"
	"os"
)

var c InterFace.ByteCounter

func main() {
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	// assert
	var w io.Writer = os.Stdout
	if _, ok := w.(*os.File); ok {

	} // success:  ok, f == os.Stdout
	if _, ok := w.(*bytes.Buffer); ok {

	} // failure: !ok, b == nil

}
