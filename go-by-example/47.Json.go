package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	// 默认json字段为Page和Fruits
	Page int
	Fruits []string
}

type Response2 struct {
	// 自定义字段，小写page, fruits
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

var pln = fmt.Println

func main() {

	// json 序列化

	// 基本类型 转 json
	bolB,_ := json.Marshal(true)
	// 转成string 进行打印输出
	pln(string(bolB))

	intB, _ := json.Marshal(1)
	pln(string(intB))

	fltB, _:= json.Marshal(1.22)
	pln(string(fltB))

	strB, _:= json.Marshal("gopher")
	pln(string(strB))


	// 数组切片等类型
	slcD := []string{"apple","banana"}
	slcB, _:= json.Marshal(slcD)
	pln(string(slcB))

	mapD := map[string]int{"first":1, "second":2}
	mapB, _ := json.Marshal(mapD)
	pln(string(mapB))

	// 结构体构建复杂的json数据
	res1D := &Response1{
		1,
		[]string{"apple", "banana"},
	}

	res1B, _ := json.Marshal(res1D)
	pln(string(res1B))

	res2D := &Response2{
		2,
		[]string{"apple", "banana"},
	}
	res2B,_ := json.Marshal(res2D)
	pln(string(res2B))


	// json反序列化
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	pln(dat)

	// 获取num
	num := dat["num"].(float64)
	pln(num)

	// 获取strs内容 断言为interface类型
	strs := dat["strs"].([]interface{})
	pln(strs[0].(string))

	// 继续json反序列化, 反序列化值结构体
	str := `{"page": 1, "fruits": ["apple", "peach"]}`

	res := &Response2{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}

	fmt.Println(res)
	// 使用结构体.属性的方式进行访问
	fmt.Println(res.Fruits[0])

	// 使用NewEncoder将json序列化后的结果输出到os.Stdout中
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}