package main

import "fmt"

//接口和继承解决问题的不同
//1. 接口: 侧重点为设计, 各种规范, 让其他自定义类型去实现这些方法.
//2. 继承: 主要解决代码复用, 和可维护性.

type Monkey struct {
	Name string
}

// 猴子自带会爬树能力

func (self Monkey) climing() {
	fmt.Println("生来会爬树")
}

// littleMonkey结构体
type littleMonkey struct {
	Monkey //继承Monkey的属性以及方法
}

type bigMonkey struct {
	Monkey //继承Monkey的属性以及方法
}

// 让littleMonkey实现flying方法.
// 1. 直接给Monkey实现flying方法,这样的话,父类被添加了功能,存在污染.
// 2. 直接给littleMonkey定义flying方法, 让littleMonkey实现. 这种方式, flying方法只会在littleMonkey上. 如果某个其他对象想要flying方法,则需要自己写一个.

func (self littleMonkey) flying() {
	fmt.Println("后期会飞")
}

func (self bigMonkey) flying() {
	fmt.Println("大猴子后期会飞")
}


// 3. 直接定义接口, 包含flying方法. 这样,就可以用多态的理念直接实现,
type iCanFly interface {
	// 声明flying()方法
	flying()
}

// 定义testMonkey结构体,
type testmonkey struct {

}
// 实现方法, flying()方法
func (self testmonkey) testFly(icanfly iCanFly) {
	icanfly.flying()
}

func main() {
	lm := littleMonkey{
		Monkey{
			Name:"wukong",
		},
	}
	lm.climing()
	lm.flying()

	bm := bigMonkey{
		Monkey{
			Name:"wuneng",
		},
	}
	tm:= testmonkey{}
	tm.testFly(bm)
}