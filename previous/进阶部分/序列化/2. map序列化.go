package main

import (
	"encoding/json"
	"fmt"
)

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})

	a["name"] = "yisan"
	a["age"] = 20
	a["addr"] = "China"
	a["gender"] = "male"

	// 进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}


func main() {
	testMap()
}

