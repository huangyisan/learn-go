package Mprint

import (
	"fmt"
	"math"
)

func FloatPrint() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f g=%[2]g \n", x, math.Exp(float64(x)))
	}
}
