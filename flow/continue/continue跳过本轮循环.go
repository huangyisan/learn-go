package main

import "fmt"
//continue 跳过本轮循环，直接到下一轮循环
func main(){
	for i:=0; i<10; i++{
		if i ==2 {
			continue
		}
		fmt.Printf("为%d\n",i)
	}
}
