package main

import "fmt"

// map的排序
// go中没有专门的方法对map的key进行排序
// go中的map也是无序的
// go中map的排序, 是先将key进行排序,然后根据key对值进行遍历.


func main() {
//	map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 10
	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	map1[4] = 4

	fmt.Println(map1)
}