package main

import (
	"fmt"
	"learn/Mprint"
)

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666" %o 八进制
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	//%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
	//Mprint.MyPrint()
	//Mprint.ChPrint()
	//Mprint.FloatPrint()
	//Mprint.ComplexPrint()
	Mprint.StringPrint()
}
