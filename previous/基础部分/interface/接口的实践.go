package main

import (
	"fmt"
	"sort"
)

// 实现对hero结构体切片的排序: sort.Sort(data interface)

// 1. 声明hero结构体
type Hero struct {
	Name string
	Age int
}

// 2. 声明hero结构体为切片类型
type HeroSlice []Hero

// 3. 实现Sort的接口三个方法
// Len() int

// i < j则按照升序排列, i>j则降序排列
//Less(i, j int) bool

// 交换i和j
//Swap(i, j int)

func (self HeroSlice) Len() int {
	return len(self)
}

//Less方法决定使用什么标准来排序, 如果按照年龄来写,从小到大排序
func (self HeroSlice) Less(i, j int) bool {
	return self[i].Age < self[j].Age
}

func (self HeroSlice) Swap(i,j int) {
	self[i] , self[j] = self[j], self[i]
}


// 对结构体hero进行排序
func main() {
	var heroes HeroSlice
	for i:=0 ;i <10 ;i++ {
		hero := Hero{
			Name:fmt.Sprintf("英雄-%d", i),
			Age:i,
		}
	//	 将hero append到heroes这个slice里
	heroes = append(heroes, hero)
	}

//	排序前的顺序
	for _,v := range heroes {
		fmt.Println(v)
	}
//	进行排序
	sort.Sort(heroes)
// 排序后的顺序
	for _,v := range heroes {
		fmt.Println(v)
	}
}