package main

// 处理者接口, 每个环节都实现这些方法

type department interface {
	execute(*patient)
	setNext(department)
}
