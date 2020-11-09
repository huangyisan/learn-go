package main

import "fmt"

type action interface {
	Say()
}

type action2 interface {
	Hello()
}

type Human struct {

}

func (s Human) Say() {
	fmt.Println("say func")
}

func (s Human) Hello() {
	fmt.Println("say func")
}

func main() {
	var human Human
	var a action = human
	var b action2 = human
	a.Say()
	b.Hello()

}