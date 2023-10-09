package main

import "fmt"

func MinAny[T ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type ordered interface {
	~int | ~float64
}

func MinAny2[T ordered, T1 string](a, b T, c T1) T {
	if a < b {
		return a
	}
	fmt.Println(c)
	return b
}

type MinSalary[T int | float64] struct {
	Salary T
}

type Salary[T int | float64] struct {
	money T
}

func (s *Salary[T]) Min(x, y T) T {
	if x < y {
		return x
	}
	fmt.Println(s.money)
	return y
}

type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

var a MyMap[string, float32] = map[string]float32{
	"j": 4.0,
	"k": 5.0,
}

type MyStruct[T int | string, T1 float32 | float64] struct {
	Name T
	Data T1
}

type IPrintData[T int | string] interface {
	Print(data T)
}

type myChain chan int
type MyChan[T int | string] chan T

type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
}

//type CommonType[T int | string | float32] T

type NewType[T interface{ *int }] []T
type NewType2[T *int,] []T
type NewType3[T interface{ *int | *float64 }] []T

type Slice[T int | string | float32 | float64] []T

//type UintSlice[T uint|uint8] Slice[T]

type WowMap[T int | string] map[string]Slice[T]
type WowMap2[T Slice[int] | Slice[string]] map[string]T

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// 队列
type Queue[T any] struct {
	elements []T
}

// 放入尾部
func (q *Queue[T]) Put(element T) {
	q.elements = append(q.elements, element)
}

// 从队列头部取出,并且从头部删除
func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, false
	}
	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) > 0
}

// 队列大小
func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func Add[T int | float64](a, b T) T {
	return a + b
}

type A struct {
}

//func (receiver A) Add[T int | float64](a, b T) T {
//	return a + b
//}

type B[T int | string | float32] struct {
}

func (receiver B[T]) Add(a, b, c T) {
	fmt.Println(a, b, c)
}

func main() {
	salary := MinSalary[int]{
		Salary: 1000,
	}
	fmt.Println(salary.Salary)

	var q1 Queue[int]
	q1.Put(1)
	a, b := q1.Pop()
	fmt.Println(a, b)

	Add[int](1, 2)

	var b1 B[int]
	b1.Add(1, 2, 3)
}

type Slice1[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64] []T

type IntUintFloat interface {
	any
}

type Slice2[T IntUintFloat] []T

func define() {
	s := Slice2[any]{1, 2, 3, "string"}
	fmt.Println(s)
}

type MyMap01[KEY comparable, VALUE any] map[KEY]VALUE
