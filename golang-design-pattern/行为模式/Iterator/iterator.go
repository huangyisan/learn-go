package main

// 迭代器要实现的方法
type iterator interface {
	hasNext() bool
	getNext() *user
}
