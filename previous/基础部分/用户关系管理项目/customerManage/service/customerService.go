package service

import "../model"

// 该结构体完成对Customer的操作
// crud行为
type CustomerService struct {
	// 存储多个customer
	customers []model.Customer
	//	声明一个字段， 表示当前切片有多少个用户
	// 该字段后面，可以作为新客户的编号id+1
	customerNum int

}

// 工厂模式返回一个NewCustomerService

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	custoer := model.NewCustomer(1,"zhangsan", "male", 20, "110", "xxx@q.com")
	customerService.customers = append(customerService.customers, custoer)

	return customerService
}

// 返回客户切片信息

func (c *CustomerService) List() []model.Customer{
	return c.customers
}

// 添加客户到customers切片中

func (c *CustomerService) Add(customer model.Customer) bool {
	// 确定一个分配id的规则， 就是添加的顺序
	c.customerNum +=1
	// NewCustomer2方法中，产生零值的Id在这里被覆写从而修复。
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}

// 根据id查找客户在切片中对应的下标， 没有，则返回-1
func (c *CustomerService) FindById(id int) int {

	// 默认为找不到, 返回-1
	index := -1
	//	遍历切片
	for i:=0; i<len(c.customers); i++ {
		if c.customers[i].Id == id {
			index = i
		}
	}
	return index

}

// 根据id删除客户

func (c *CustomerService) DeleteById (id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}

	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

// 修改用户信息

func (c *CustomerService) ModifyById (id int, name string, gender string, age int, phone string, email string) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	if name == "" {
		name = c.customers[index].Name
	}

	if gender == "" {
		gender = c.customers[index].Gender
	}

	if age == -1 {
		age = c.customers[index].Age
	}

	if phone == "" {
		phone = c.customers[index].Phone
	}

	if email == "" {
		email = c.customers[index].Email
	}

	c.customers[index].Name = name
	c.customers[index].Gender = gender
	c.customers[index].Age = age
	c.customers[index].Phone = phone
	c.customers[index].Email = email
	return true
}