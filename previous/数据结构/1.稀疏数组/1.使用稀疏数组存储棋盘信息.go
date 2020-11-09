package main

import "fmt"

// 7*7棋盘
// 数字1表示白旗
// 数字2表示黑旗
// 使用稀疏数组来表示整个棋盘布局

/*
0 0 0 0 0 0 0
0 1 0 0 0 0 0
0 0 2 0 1 0 0
0 0 0 0 0 0 0
0 0 0 0 0 0 0
0 0 0 2 0 0 0
0 0 0 0 0 0 0
 */

//
func main() {
	// 定义一个7*7的二维数组, 将黑子白子进行赋值.
	var chessMap [7][7]int
	chessMap[1][1] = 1
	chessMap[2][2] = 2
	chessMap[2][4] = 1
	chessMap[5][3] = 2

	// 打印原始数据
	for _,v := range chessMap {
		for _,v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Printf("\n")
	}

	// 用稀疏数组表示
	/* 每一个棋子都有三个属性
		1. x轴位置
		2. y轴位置
		3. 棋子颜色
		这三个属性可以使用结构体来表示
		然后将结构体依次append进数组中, 就可以作为一个稀疏数组了.
	 */

	 // 定义一个存放棋子属性的结构体
	 type chess struct {
	 	axisX int
	 	axisY int
	 	value int
	 }

	 // 定义一个存放棋子结构体的切片
	 var chessSparseMap []chess

	 // chessSparseMap稀疏数组第一行定义x,y和value
	 firstLine := chess{
	 	axisX: 7,
	 	axisY: 7,
	 	value: 0,
	 }

	 chessSparseMap = append(chessSparseMap, firstLine)

	 // 对原始数据进行遍历, 然后将获取到有数值内容的值,存储到chess结构体中, 并且将该结构体append到chessSparseMap里面.
	for k,v := range chessMap {
		for k2,v2 := range v {
			if v2 != 0 {
				tmpChess := chess{
					axisX:k,
					axisY:k2,
					value:v2,
				}
				chessSparseMap = append(chessSparseMap, tmpChess)
			}
		}
	}
	fmt.Println("稀疏数组为")
	fmt.Println(chessSparseMap)

}