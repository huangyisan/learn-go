package main

// 集合接口
type collection interface {
	createIterator() iterator
}
