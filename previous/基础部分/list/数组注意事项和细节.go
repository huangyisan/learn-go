package main

import "fmt"

/*
1. 数组是多个相同类型数据的组合, 一个数组一旦声明/定义后, 其长度是固定的,不能改变.
2. 如果没有定义长度, 则该数组为切片. 不定长度[...]还是属于数组.
3. 数组中的值可以是任何类型, 可以是引用类型, 也可以是值类型, 但类型不可以混用.
4. 数组创建后, 如果没有赋值, 则默认为对应类型的零值.
5. 使用数组步骤, 1. 声明数组开辟内存空间, 2. 给数组各个元素赋值, 3. 使用数组.
6. 数组的下标是从0开始的.
7. 数组下标必须在范围内, 否则报panic 越界错误.
8. go的数组属于值类型, 因此会进行值拷贝, 数组间不会相互影响.
9. 如果想在其他函数中, 去修改原来的数组, 可以使用引用传递(指针方式)
 */

 //第八点测试
 func test(arr [4]int) {
 	arr[0] = 0
 	fmt.Printf("我是test函数中的arrlist%v\n", arr)
 }
 // 使用指针传递的方式,修改原来的数组,第九点测试
 func test1(arr *[4]int) {
 	 arr[0] = 0
 	//  也可以写成以下写法, 等同上面arr[0] = 0写法. 数组中(*arr)[0]可以简写为arr[0], 但切片中不可以.
	 //(*arr)[0] = 0
	 fmt.Printf("我是test1函数中的arrlist%v\n", arr)

 }

 func main() {
 	arr := [4]int{1,2,3,4}

 	//test(arr)
	 //// 因为是值拷贝, 所以test中对arr的修改并不会影响arr的值.
 	//fmt.Println("我是main函数中的arrlist", arr)

 	//使用指针传递的方式,修改原来的数组.
 	test1(&arr)
 	fmt.Println("我是main函数中的arrlist", arr)

 }