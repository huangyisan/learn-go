package functions

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func CallHypot() {
	fmt.Println(hypot(3, 4))
}

func changeSlice(slice []string) {
	slice[0] = "abc"

}

func CallchangeSlice() {
	slice := []string{"a", "b", "c"}
	changeSlice(slice)
	fmt.Println(slice)
}
