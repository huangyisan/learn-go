package main

import (
    "flag"
    "fmt"
    "os"
)

// 使用flag处理子命令
// 比如go build和go get是两个不同的命令，后面跟随的参数也不一样。

func main() {
    // 使用NewFlagSet函数声明一个子命令
    fooCmd := flag.NewFlagSet("foo",flag.ExitOnError)
    fooEnable := fooCmd.Bool("enable", true, "enable foo")

    fooName := fooCmd.String("fooname","im foo", "set foo name")

    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    barLevel := barCmd.Int("level", 0, "set bar level")

    // 当不传入任何参数的时候，长度为1，为执行程序绝对路径
    if len(os.Args) == 1 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    // 当选择不同的子命令的时候
    switch os.Args[1] {
    case "foo":
        // parse foo的子命令
        fooCmd.Parse(os.Args[2:])
        fmt.Println("SubCommand 'foo'")
        fmt.Println("enable:", *fooEnable)
        fmt.Println("fooname:", *fooName)
        fmt.Println("args:", fooCmd.Args())

    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("BarCommand 'foo'")
        fmt.Println("barLevel: ", *barLevel)
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}


/*
../go-by-example (master)
$ ./66.Command-Line\ Subcommands.exe -foo -h
expected 'foo' or 'bar' subcommands

../go-by-example (master)
$ ./66.Command-Line\ Subcommands.exe foo -h
Usage of foo:
  -enable
        enable foo (default true)
  -fooname string
        set foo name (default "im foo")

../go-by-example (master)
$ ./66.Command-Line\ Subcommands.exe foo -enable false -fooname myfo
SubCommand 'foo'
enable: true
fooname: im foo
args: [false -fooname myfo]

 */