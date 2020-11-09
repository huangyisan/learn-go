package model

import "fmt"

type person struct {
	Name string
	age int
	salary float32
}

func NewPersion(name string, age int, salary float32) *person {
	return &person{
		Name:name,
		age:age,
		salary:salary,
	}
}

func (p *person) SetAge(age int) {
	if age <= 0{
		fmt.Println("输入年龄异常")
	} else {
		p.age = age
	}
}

func (p *person) GetAge() int{
	return p.age
}