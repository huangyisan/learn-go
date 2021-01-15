package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 将标准输入作为参数，创建衣蛾scanner实例
	scanner := bufio.NewScanner(os.Stdin)

	// Scan返回bool，如果读完则返回false，反之返回true
	for scanner.Scan() {
		// Text()返回字符串
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(scanner.Text())
		fmt.Println(ucl)
	}

	// 如果存在err，则将err打印输出到 错误输出里
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}