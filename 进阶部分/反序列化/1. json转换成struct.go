package main

import (
	"encoding/json"
	"fmt"
)

// 根据json字符串内容写出struct
type Monster struct {
	// 此时json转出来的Name为大写，如果要小写，或者定义为其他的，在可以给他加个json的tag，这样json序列化后就采用tag中的名称。
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string `json:"shenri"`
	Sal float64 `json:"sal"`
	Skill string `json:"skill"`
}

func convStruct() {

	str := "{\"name\":\"cow\",\"Age\":100,\"shenri\":\"2010-01-01\",\"Sal\":8000,\"Skill\":\"punch\"}"

	var monster Monster
	// 反序列的字符串，传递给monster. 返回为错误
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println(err)
	}

	// 反序列化后可以获取到对应的内容
	fmt.Println(monster.Birthday, monster.Name)

}

func main() {
	convStruct()
}


