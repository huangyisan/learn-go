package utils

// 为了让其他文件import， 函数名称首字母必须大写
func Cal(n1 int, n2 int, op byte) int {
	var res int
	switch op {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	}
	return res
}
