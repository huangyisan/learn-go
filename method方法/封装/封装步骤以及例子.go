package main

import (
	"./model"
	"fmt"
)

// 1. 将结构体、字段（属性）的首字母小写（不让导出）。
// 2. 将结构体所在的包提供一个工厂函数，首字母大写，类似一个构造函数。
// 3. 提供一个首字母写的Set方法，对属性进行判断和赋值。
// 4. 提供一个首字母写的Get方法，用来获取属性的值。

// 例子
// 不能随便查看年龄、工资，对输入的年龄进行校验和判断。

func main() {
	var age int
	Lilei := model.NewPersion("lilei",20, 2000)
	fmt.Println("修改年龄:")
	fmt.Scanf("%d", &age)
	Lilei.SetAge(age)
	fmt.Printf("%s当前年龄为%d", Lilei.Name,Lilei.GetAge())
}