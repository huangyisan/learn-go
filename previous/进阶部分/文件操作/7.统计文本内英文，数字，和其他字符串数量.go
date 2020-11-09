package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

// 结构体用于保存统计结果

type CharCount struct {
	ChCount int // 记录英文个数
	NumCount int
	SpaceCount int // 空格个数
	OtherCount int // 其他字符个数
}


func main() {
	// 定义一个charCount实例
	var charCount CharCount
	// 打开一个文件，创建reader，每读取一行， 则统计改行有多少个英文，数字，空格和其他字符。
	// 将结果保存到结构体中
	filePath := "C:\\Users\\hysan\\Desktop\\github\\hys\\learn-go\\进阶部分\\文件操作\\book"
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	newReader := bufio.NewReader(file)

	for {
		str, err := newReader.ReadString('\n')

		for _,v := range str {
			// v得到的为ascii码， 可以直接和字符作比较。
			if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
				charCount.ChCount += 1
			} else if v >= '0' && v <= '9' {
				charCount.NumCount += 1
			} else if v == 32 {
				charCount.SpaceCount += 1
			} else {
				charCount.OtherCount += 1
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("英文字母有%d个\n数字有%d个\n空格有%d个\n其他字符有%d个\n", charCount.ChCount,
		charCount.NumCount, charCount.SpaceCount, charCount.OtherCount)

}