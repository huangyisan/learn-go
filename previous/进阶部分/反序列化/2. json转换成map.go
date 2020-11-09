package main

import (
	"encoding/json"
	"fmt"
)

func convMap() {
	str := "{\"addr\":\"China\",\"age\":20,\"gender\":\"male\",\"name\":\"yisan\"}"
	// 根据json字符串申明map
	var a map[string]interface{}

	// 虽然map是引用类型，但这边反序列化需要传入map的地址，也就是&map
	// 反序列化map的时候，不需要make， 因为make的操作被封装到Unmarshal函数里面了
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func main() {
	convMap()
}