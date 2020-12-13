package main

import "fmt"

// 方法特征的集合为interface接口.

type caculate interface {
	area() int
	long() int
}

type rectangle struct {
	x int
	y int
}

type strange struct {
	x int
	y int
}

func (r rectangle) area() int {
	return r.x * r.y
}

func (r rectangle) long() int {
	return (r.x + r.y) * 2
}

func (s strange) area() int {
	return (s.x + s.y) * s.x * s.y
}

func (s strange) long() int {
	return (s.x + s.y) * 3
}


func calc(c caculate)  {
	fmt.Println(c.area())
	fmt.Println(c.long())
}

func main() {
	cfx := rectangle{
		3,
		4,
	}
	// 根据传入的结构体不同,得到的结果也不同
	calc(cfx)

	str := strange{
		3,
		4,
	}

	calc(str)
}
