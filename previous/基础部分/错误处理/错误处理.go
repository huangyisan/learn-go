package main

import "fmt"

// 1. 默认情况下，当发生错误panic 程序就会崩溃退出
// 2. 如果我们希望，当错误发生后， 可以捕获到错误，并进行处理，保证代码可以继续向下执行。还可以捕获到错误后，作出一些动作。
// 3. 错误处理方式 defer, panic, recover.
// 4. 处理过程， go中抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
// 5. recover得放在defer中使用

func test() {
	//使用defer + recover 来捕获和处理异常

	defer func() {
		// recover() 是内置函数，可以捕获到异常
		err := recover()
		// 如果err不为nil， nil是err的零值，如果不等于，则表示err有内容，有错误，捕获到了异常
		if err != nil {
			// 如果出现错误，则执行里面内容
			fmt.Println(err)
		}
	//	直接调用该匿名函数
	}()

	num1 := 1
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func main() {
	test()
	// 当对错误进行处理后，后面代码可以正常执行
	fmt.Println("抛出错误后的代码")
}