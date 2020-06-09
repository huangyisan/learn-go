package main

import "fmt"


/*
1. 定义两个标记, head和tail.
2. head表示队列头部(但不包含队首元素, head随着元素的弹出而改变), tail表示队列尾部(tail随着元素的推入而改变).
3. 当tail等于head, 此时表示队列为空.
4. 当tail等于队列容量-1, 此时表示队列已满(之所以-1,是因为tail是从-1开始).
5. tail从-1开始,每次塞入一个数据,tail+1
6. head从-1开始,每次弹出一个数据, head+1
 */

// 声明队列结构体, 四个要素, 长度, head, tail, 数组
type queue struct {
	head int
	tail int
	array [4]int
	maxCap int
}


func (q *queue) pushQueue(value int) {
	// 先判断队列是否满
	if q.tail == q.maxCap {
		fmt.Println("队列已满")
		return
	}

	q.array[q.tail] = value
	q.tail +=1

}

func (q *queue) popQueue() {
	if q.tail == q.head {
		fmt.Println("队列为空")
		return
	}

	val := q.array[q.head]
	q.head +=1
	fmt.Println(val)

}

func (q *queue) showQueue() {
	for i:=q.head +1 ; i<q.tail; i++ {
		fmt.Println(q.array[i])
	}
}