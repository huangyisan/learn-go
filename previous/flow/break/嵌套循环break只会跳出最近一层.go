package main

import "fmt"
// 多层嵌套， break只会跳出最近一层
func main(){
	var i,j int
	for i=0;i<200; i++ {
		for j=0; j<10;j++{
			if j==9 {
				fmt.Println("执行break")
				break
			}
		}

		fmt.Printf("i为%d\n", i)
	}
}