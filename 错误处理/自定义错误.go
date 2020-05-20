package main
// 支持自定义错误，使用errors.New和panic内置函数。
// 1. errors.New("错误说明")， 会返回一个error类型的值，表示一个错误。
// 2. panic内置函数，接受一个interface{}类型的值（任何值）作为参数。可以接受error类型的变量，输出错误信息，并退出程序。

