package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个struct

type Monster struct {
	// 此时json转出来的Name为大写，如果要小写，或者定义为其他的，在可以给他加个json的tag，这样json序列化后就采用tag中的名称。
	Name string `json:"name"`
	Age int
	Birthday string `json:"shenri"`
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
