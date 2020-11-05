package main

import (
	"encoding/json"
	"fmt"
)

func convSlice() {
	str := "[{\"addr\":\"China\",\"age\":28,\"cc\":\"bb\",\"name\":\"yisan\"}," +
		"{\"addr\":\"UK\",\"age\":19,\"bb\":[\"uu\",\"kk\"],\"name\":\"lilei\"}]"

	var slice []map[string]interface{}
	// 不需要对slice进行make， 因为Unmarshal里面已经封装了make。
	// 传入的是地址， &slice
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice)

}

func main() {
	convSlice()
}
