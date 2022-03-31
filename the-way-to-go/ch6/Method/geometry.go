package Method

import "math"

type Point struct {
	X, Y float64
}

// function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// 使用指针的方式, 函数的参数默认是值拷贝,如果函数中要变更一个变量,或者参数太大,默认要改变这种方式,则可以传入指针.
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
