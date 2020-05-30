package stuPackage

// student小写无法导出被main函数使用
type student struct {
	Name string
	age int
}

// 通过工厂模式来解决
// 定义一个方法， 返回指针， 该指针用来指向student， 这样调用这个函数，等于是使用了student这个结构体了。
// 当然这个函数名称首字母得大写。

func NewStudent(n string, a int) *student {
	return &student{
		Name:n,
		age:a,
	}
}

// 获取给student结构体绑定一个GetAge的方法， return 结构体的age， 然后main函数进行调用该方法。
func (s *student) GetAge() int {
	return s.age
}