package main

import (
    "fmt"
    "testing"
)

// 以Test开头之后跟随第一位大写，作为测试函数
func TestIntMin(t *testing.T) {
    ans := IntMin(1,2)
    if ans != 1 {
        // t.Error*报错后还会继续执行
        t.Errorf("IntMin(1,2) result %d, want 1", ans)
    }
}

// 当存在很多数据的时候，使用表驱动的方式进行测试
func TestTableIntMin(t *testing.T)  {
    var tests = []struct{
        a int
        b int
        want int
    } {
        // 待测试数据
        {1,2 ,1},
        {-11,2 ,-11},
        {-1,-2 ,-2},
        {0,1 ,0},
    }

    for _, v := range tests {
        testname := fmt.Sprintf("%d, %d, want %d", v.a, v.b, v.want)
        // t.Run 可以运行一个 “subtests” 子测试，一个子测试对应表中一行数据。 运行 go test -v 时，他们会分开显示。
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(v.a, v.b)
            if ans != v.want {
                t.Errorf("got %d, want %d", ans, v.want)
            }
        })
    }
}

// 进入该目录，执行go test -v