package main

import "fmt"

type tv struct {
	state bool
}

func (t *tv) on() {
	t.state = true
	fmt.Println("tv on")
}

func (t *tv) off() {
	t.state = false
	fmt.Println("tv off")
}
