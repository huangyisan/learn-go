package ch1

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	fmt.Print(123)
	t.Logf("%v", Foo(1, 2))
}
