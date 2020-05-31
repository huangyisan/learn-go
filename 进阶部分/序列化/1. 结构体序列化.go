package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个struct

type Monster struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

func testStruct() {
	monster := Monster{
		Name:"cow",
		Age:100,
		Birthday:"2010-01-01",
		Sal: 8000.0,
		Skill: "punch",
	}

	// 序列化成json
	data, err := json.Marshal(&monster)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}


func main() {
	testStruct()
}
