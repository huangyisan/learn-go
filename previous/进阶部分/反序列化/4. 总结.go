package main

// 切片， 结构体，map的定义需要和传递方约定好。确保序列化前和序列化后的类型一致。
// 对于结构体，一般可能会用到json tag， 因为结构体字段可能为大写，而json的字段为小写。
// 反序列化传入的是地址。
// 对于map和切片，不需要make，因为Unmarshel里面已经封装好了。