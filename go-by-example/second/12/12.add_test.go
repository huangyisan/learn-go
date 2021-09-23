package _2

import "testing"

func TestAdd(t *testing.T) {
	if ret := add(1,2); ret != 3 {
		t.Error("add 1 2 should be 3")
	}
}
