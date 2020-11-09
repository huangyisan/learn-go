package main

import (
	"encoding/json"
	"fmt"
)

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "yisan"
	m1["age"] = 28
	m1["addr"] = "China"
	m1["cc"] = "bb"

	var m2 = make(map[string]interface{})
	m2["name"] = "lilei"
	m2["age"] = 19
	m2["addr"] = "UK"
	m2["bb"] = []string{"uu","kk"}

	slice = append(slice, m1,m2)
	slice[0] = m1
	slice[1] = m2


	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}


func main() {
	testSlice()
}