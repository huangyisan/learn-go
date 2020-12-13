package main

import "fmt"

// https://studygolang.com/articles/12266#commentForm

// 根据公司员工的个人薪资，计算公司一年的总支出



// 工资算法接口，需要实现工资算法函数, mon为月份

type SalaryAlgorithm interface {
	salaryAlg(mon int) int
}

// 试用期员工和正式员工的结构体

type probationStaff struct {
	salary float64
}

// 正式员工
type regularEmployee struct {
	salary float64
}

// 定义试用期员工的工资算法

func (p probationStaff) salaryAlg(mon int) int {
	perMon := p.salary * 0.8
	total := int(perMon) * mon
	return total
}

// 正式员工

func (p regularEmployee) salaryAlg(mon int) int {
	perMon := p.salary
	total := int(perMon) * mon
	return total
}

// 定义一个求总和的函数
// 因为probationStaff和regularEmployee 都实现了salaryAlg， 所以他们都是SalaryAlgorithm
// 把SalaryAlgorithm切片作为参数传入
func totalPay(s []SalaryAlgorithm) int {
	// 迭代SalaryAlgorithm切片，然后求和
	total := 0
	for _,v := range s{
		total += v.salaryAlg(12)
	}
	return total
}

func main() {
	// 初始化2个 试用期员工，1个正式员工
	p1 := probationStaff{
		3000,
	}

	p2 := probationStaff{
		4000,
	}

	r1 := regularEmployee{
		8000,
	}

	s := []SalaryAlgorithm{p1,p2,r1}

	total := totalPay(s)

	fmt.Println(total)

}

