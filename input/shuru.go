package main

import "fmt"

func main() {
	var name string
	//var age byte
	//var sal float32

	fmt.Println("输入姓名等信息")
	fmt.Scanln(&name)
	fmt.Println(name)
	fmt.Scanf("%s",&name)
	fmt.Println(name)


}
