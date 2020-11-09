package model

import "fmt"

// 声明一个customer结构体， 表示一个客户信息

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

// 编写一个工厂模式，返回一个customer实例

func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id: id,
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}

// 编写一个不带id的工厂函数， 返回一个customer实例
// 这边返回不带id， 并不表示Customer却了Id这个变量，如果不传，那么Id为零值。
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}


// 方法显示该用户的信息
func (c* Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}