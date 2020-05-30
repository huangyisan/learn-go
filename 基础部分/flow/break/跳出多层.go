package main
// 使用label， break指定label跳出多层

import "fmt"

func main(){
	var i,j int
	I:
		for i=0;i<200; i++ {
			fmt.Printf("i为%d\n", i)
			for j=0; j<10;j++{
				if j==9 {
					fmt.Println("执行break")
					//指定跳出到I
					break I
				}
			}
		}
}
