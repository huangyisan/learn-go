package main
// 该文件为cal_test.go 命名也是有规则的，需要xx_test.go
// testing 为go的轻量级框架包
import (
	"testing"
)

// 编写一个测试用例，来测试addUpper是否正确

// 测试函数命名法则： 函数定义必须为Test开头，后面跟随的第一个字符为非a-z的小写字母。用来识别测试程序
func TestAddUpper(t *testing.T) {

	// 调用
	res := addUpper(10)
	// 使用Fatalf方法格式化输出错误信息，并且退出程序
	if res != 55 {
		t.Fatalf("addUpper执行错误，期望值为%v,实际值为%v", 45, res)
	}

	// 正确，则输出日志
	t.Logf("执行正确")
}


// 在控制台使用go test -v进行测试
/*
$ go test -v
=== RUN   TestAddUpper
--- PASS: TestAddUpper (0.00s)
    cal_test.go:21: 执行正确
PASS
ok      _/C_/Users/hysan/Desktop/github/hys/learn-go/进阶部分/单元测试/testCase10.038s

 */