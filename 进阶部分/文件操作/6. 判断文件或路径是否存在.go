package main
// golang使用os.Stat()函数判断文件或者路径是否存在.
// 如果返回值为nil表示存在
// 如果错误类型使用os.IsNotExist()判断为true,表示不存在
// 如果返回错误为其他类型,则不能确定是否存在.
