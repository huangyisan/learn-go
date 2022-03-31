package Method

import (
	"fmt"
	"image/color"
)

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func PrintColor() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.X)

	red := color.RGBA{255, 0, 0, 255}
	var p = ColoredPoint{Point{1, 2}, red}
	// ColoredPoint嵌套了匿名Point,扩展了ColoredPoint的方法,继承了Point的caleBy
	p.ScaleBy(2)
}
