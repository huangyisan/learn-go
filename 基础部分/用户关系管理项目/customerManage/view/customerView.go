package main

import (
	"fmt"
	"../service"
	"../model"
)


type customerView struct {
	//	接受用户的输入
	key string
	// 是否循环显示
	loop bool

//	增加一个customerService字段
	customerService *service.CustomerService

}


// 显示客户信息

func (c *customerView) list() {
//	 获取当前所有客户信息
	customers := c.customerService.List()

	fmt.Println("-----客户列表------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i:=0; i<len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
}

// 修改用户方法
func (c *customerView) Modify() {
	fmt.Println("-----修改客户-----")
	fmt.Println("输入待修改id")
	id := -1
	fmt.Scanln(&id)
	fmt.Println("输入修改姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("输入修改性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("输入修改年龄:")
	age := -1
	fmt.Scanln(&age)
	fmt.Println("输入修改手机:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("输入修改邮箱:")
	email := ""
	fmt.Scanln(&email)

	if c.customerService.ModifyById(id, name, gender, age, phone, email) {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}





}


// 添加用户方法
func (c *customerView) add() {
	fmt.Println("-----添加客户-----")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	// 将得到的信息进行初始化成一个用户信息
	//	 id为系统分配

	customer := model.NewCustomer2(name, gender, age, phone, email)
	if c.customerService.Add(customer) {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}

}

// 删除用户方法
func (c *customerView) delete() {
	fmt.Println("------删除用户-------")
	fmt.Println("请输入删除客户编号：")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	fmt.Printf("是否确认删除ID为%v的用户？[N/Y]", id)
	var enter string
	flag := false
	for {
	fmt.Scanln(&enter)
	switch enter {
		case "Y" :
			if c.customerService.DeleteById(id) {
				fmt.Println("删除成功")
			} else {
				fmt.Println("删除失败")
			}
			flag = true
		case "N":
			fmt.Println("取消删除")
			flag = true
		default:
			fmt.Println("请输入N或者Y")


		}
		if flag {
			break
		}
	}
	//if flag == "Y" {
	//	if c.customerService.DeleteById(id) {
	//		fmt.Println("删除成功")
	//	} else {
	//		fmt.Println("删除失败")
	//	}
	//} else {
	//	return
	//}


}


// 显示主菜单

func (c *customerView) mainMenu() {
	for {
		fmt.Println("------客户信息管理系统------")
		fmt.Println("      1. 添加用户")
		fmt.Println("      2. 修改用户")
		fmt.Println("      3. 删除用户")
		fmt.Println("      4. 用户列表")
		fmt.Println("      5. 退    出")
		fmt.Println("      请选择(1-5):")
		
		fmt.Scanln(&c.key)

		switch c.key {
			case "1":
				//fmt.Println("添加客户")
				c.add()
			case "2":
				//fmt.Println("修改客户")
				c.Modify()
			case "3":
				//fmt.Println("删除客户")
				c.delete()
			case "4":
				//fmt.Println("用户列表")
				c.list()
			case "5":
				c.loop = false
			default:
				fmt.Println("请输入正确编号")
		}
		if !c.loop {
			break
		}

	}
	fmt.Println("已经退出系统")
}


func main() {
//	在main函数中， 创建实例
	customerview := customerView{
		key: "",
		loop: true,
	}
	// 初始化customerService字段, 将用户信息初始化
	customerview.customerService = service.NewCustomerService()

	customerview.mainMenu()

}