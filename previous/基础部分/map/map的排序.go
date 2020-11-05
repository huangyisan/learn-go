package main

import (
	"fmt"
	"sort"
)

// map的排序
// go中没有专门的方法对map的key进行排序
// go中的map也是无序的. 之前1.10版本生成的map 不会按照key自动排序, 升级到1.14后,发现已经是按照key自动排序了.
// go中map的排序, 是先将key进行排序,然后根据key对值进行遍历.


func main() {
//	map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 10
	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	map1[4] = 4

	// 将key放入切片, 然后对切片排序, 再遍历切片, 获取key, 按照key的顺序输出map

	var keys []int

	for k,_ := range map1 {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Ints(keys)

	// 根据排序后的keys 取出value, 得到map
	for _, i := range keys {
		fmt.Printf("map[%v]=%v \n", i, map1[i])
	}

	fmt.Println(map1)
}