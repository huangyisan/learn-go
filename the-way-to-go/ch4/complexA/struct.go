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
