package main

import (
	"math/rand"
	"fmt"
	"time"
)
/*
思路
1. 使用rand.Intn获取随机数, 但如果只用rand.Intn则每次生成的随机数都一样,需要使用seed进行乱数.
2. 将随机数放入一个数组中
3. 反转数组? 交换次数刚好是数组长度的一半
 */

func main() {
	// 定义一个长度为5的数组
	var intArry [5]int

	// 为了每次的seed值不一样, 使用时间戳进行改变seed.
	rand.Seed(time.Now().Unix())

	// 获得数组长度的一般.
	halfLength := len(intArry) / 2

	// 将随机数放入数组中
	for i:=0; i< len(intArry); i++ {
		intArry[i] = rand.Intn(100) //0<=x<100, 无法取到100
	}
	fmt.Println("原先数组为",intArry)
	//进行交换,交换次数为长度的一半,向下取整.

	for l:=0; l < halfLength;l++ {
		intArry[l], intArry[len(intArry)-1-l] = intArry[len(intArry)-1-l], intArry[l]

	}

	fmt.Println("交换后数组为",intArry)
}