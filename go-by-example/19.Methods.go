package main

import "fmt"

type Person struct {
	Name string
	Age int
	Addr string
}

func (P Person) getInfo(info string) interface{}  {
	switch info {
	case "Name":
		return P.Name
	case "Age":
		return P.Age
	case "Addr":
		return P.Addr
	default:
		return "unknow"
	}
}

func main() {
	p := &Person{
		"json",
		19,
		"TW",
	}
	pInfo := p.getInfo("Name")
	fmt.Println(pInfo)
}