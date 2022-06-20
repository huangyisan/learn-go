package main

import "fmt"

// 路径包含了名称, 同时路径下可能包含了文件,也可能包含了其他路径, 两者都实现inode接口,所以这里直接写inode的数组
type folder struct {
	name     string
	children []inode
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {

}
