package main

//1. args是slice，可以通过args[index]可以访问到各个值

// args...表示多个不定参数
func sum(args... int) int{
	return args[0] + args[1]

}

// 支持1到多个可变参数。 不可变参数必须在可变参数前面
func sum2(n1 int, args... int) int {
	return args[0] + args[1] + n1
}
