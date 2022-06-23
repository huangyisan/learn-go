package main

// 创建接口,兼容原始upstream和nginx
type server interface {
	handleRequest(string, string) (int, string)
}
