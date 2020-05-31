package main

import (
	"flag"
	"fmt"
)

// 使用flag包可以解析参数，且参数顺序可以随意。

func main() {
	var user string
	var pwd string
	var host string
	var port int
	// "u" 表示 -u, 第二个参数为默认值，第三个参数为说明
	flag.StringVar(&user, "u", "", "-u 指定user")
	flag.StringVar(&pwd, "p", "", "-p 指定password")
	flag.StringVar(&host, "h", "", "-h 指定host")
	flag.IntVar(&port, "P", 80, "-P 指定port")

	// 调用Parse，从os.Args[1:]中解析出flag
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
	// go run 2.use_flag_to_prase_args.go -u yisan -p sdfsdf -P 3306 -h 127.0.0.1
}