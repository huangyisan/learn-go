package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)

// 数字1表示白旗
// 数字2表示黑旗
// 给出稀疏数组为如下格式, 结构体内分别表示x,y,value三个值,进行还原成原先棋盘布局.
// 第一个元素表示棋盘区域范围
// 这边有个问题,就是被读取的文件最后一行必须有个换行.

/*
[{7 7 0} {1 1 1} {2 2 2} {2 4 1} {5 3 2}]

7 7 0
1 1 1
2 2 2
2 4 1
5 3 2
 */

//
func main() {
	// 定义单个棋子的属性, x, y, value
	type chess struct {
		axisX int
		axisY int
		value int
	}

	// 定义一个7*7数组
	var chessMap [7][7]int

	// 读取sparseMap文件
	file, err := os.Open("/Users/huangyisan/Desktop/github/hys/learn-go/数据结构/1.稀疏数组/sparseMap")
	defer file.Close()
	if err != nil {
		fmt.Println("打开异常")
		return
	}

	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("内容读取完毕")
			break
		}

		if err == io.EOF {
			break
		}
		contentSlice := strings.Split(content," ")

		axisX, err := strconv.Atoi(contentSlice[0])
		if err != nil {
			fmt.Println("axisX转换出错")
			return
		}

		axisY, err := strconv.Atoi(contentSlice[1])
		if err != nil {
			fmt.Println("axisY转换出错")
			return
		}
		// strings.Split(contentSlice[2], "\n")[0] value结尾存在\n, 通过Split方法去除\n
		value, err := strconv.Atoi(strings.Split(contentSlice[2], "\n")[0])
		if err != nil {
			fmt.Println("value转换出错",err)
			return
		}

		if value != 0 {
			chessMap[axisX][axisY] = value
		}
	}

	for _,v := range chessMap {
		for _, v2 :=range v{
			fmt.Printf("%d \t",v2)
		}
		fmt.Println()
	}
}
