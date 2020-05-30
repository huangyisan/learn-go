package main

import (
	"fmt"
	"go/types"
	"reflect"
)

func typeJudge(item ...interface{}) {
	for _, value := range item {

		switch value.(type) {
		case int:
			fmt.Println("int类型")
		case string:
			fmt.Println("字符串类型")
		case [3]int:
			fmt.Println("数组类型")
		case types.Array:
			fmt.Println("数组类型")
			
		}
	}
}

func main() {
	var myList [3]int

	
	typeJudge(myList,1)
	t := reflect.TypeOf(myList)
	fmt.Println(t)



}