package main

type inode interface {
	print(string)
	// inode 接口实现了print, clone方法,并且clone可以产生一个符合inode接口的对象
	clone() inode
}
