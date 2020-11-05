package main

import "fmt"

// 切片的数据类型如果是map, 则为map切片, 这样map的个数就可以动态变化了

// []map[string]string

func main() {
//	1. 什么一个map切片

	var monster []map[string]string

	monster = make([]map[string]string, 2)

// 2. 增加一个妖怪信息
	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "cow"
		monster[1] = make(map[string]string)
		monster[1]["name"] = "rabbit"

	//	因为最大长度make的时候为2, 想要继续往里面添加则使用append方式,否则会出现越界报错.
		newMonster := make(map[string]string)
		newMonster["name"] = "bear"
		monster = append(monster, newMonster)
	}

	fmt.Println(monster)
}