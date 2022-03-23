package complexA

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func StructDo() {
	dilbert.Salary -= 5000
	pos := &dilbert.Position
	*pos = "Senior" + *pos

	var employeeOfTheMonth *Employee
	employeeOfTheMonth.Position += "(proactive team player)"
}

type Point struct {
	X, Y int
}

func Pstruct() {
	p := Point{1, 2}
	x := Point{Y: 1, X: 3}
	fmt.Println(p)
	fmt.Println(x)
}

func CompareStruct() {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Printf("%t, %t", p == q, p.Y == q.Y)
}

type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

func CreateWheel() {
	var w Wheel
	w.Y = 8
	w.X = 8
	w.Radius = 5
	w.Spokes = 20
}

type Point01 struct {
	X, Y int
}

type Circle01 struct {
	Center Point
	Radius int
}

type Wheel01 struct {
	Circle Circle01
	Spokes int
}

func CreateWheel01() {
	var w Wheel01
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}

type Circle02 struct {
	Point
	Radius int
}

type Wheel03 struct {
	Circle02
	Spokes int
}

func CreateWheel03() {
	var w Wheel03
	w.Y = 1
	w.Y = 1
	w.Spokes = 20
	// 如果是字面量方式，必须写清楚
	w = Wheel03{Spokes: 20, Circle02: Circle02{Point: Point{Y: 2, X: 1}, Radius: 20}}
}
