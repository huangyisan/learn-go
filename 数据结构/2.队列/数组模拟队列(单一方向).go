package main

import (
	"fmt"
	"errors"
)


/*
1. 定义两个标记, head和tail.
2. head表示队列头部(但不包含队首元素, head随着元素的弹出而改变), tail表示队列尾部(tail随着元素的推入而改变).
3. 当tail等于head, 此时表示队列为空.
4. 当tail等于队列容量-1, 此时表示队列已满(之所以-1,是因为tail是从-1开始).
5. tail从-1开始,每次塞入一个数据,先tail+1, 然后塞入
6. head从-1开始,每次弹出一个数据, 先head+1, 然后弹出
 */

// 声明队列结构体, 四个要素, 长度, head, tail, 数组
type queue struct {
	head int
	tail int
	array [4]int
	maxCap int
}


func (q *queue) pushQueue(value int) (error) {
	// 先判断队列是否满
	if q.tail == q.maxCap -1 {
		fmt.Println("队列已满")
		return errors.New("队列满了")
	}

	// tail起始位置是-1, 先挪在填充
	q.tail+=1
	q.array[q.tail] = value
	return nil


}

func (q *queue) popQueue() error{
	// 弹出之前先判断数组是否为空
	if q.tail == q.head {
		fmt.Println("队列为空")
		return errors.New("队列为空")
	}

	// head永远指向队首, 但永远不是指向第一个队首元素,如果第一个队首元素是角标为2, 那么head就是1
	// 所以如果要取出真正的第一个值,需要让head先+1
	q.head +=1
	val := q.array[q.head]
	fmt.Printf("弹出了%v\n", val)
	return nil

}

func (q *queue) showQueue() {
	// head+1 开始, 然后追赶tail角标, 中间追赶部分的元素,就是队列内的元素.
	for i:=q.head +1 ; i<=q.tail; i++ {
		fmt.Println(q.array[i])
	}
}

func main() {
	Q := &queue{
		head:-1,
		tail:-1,
		maxCap:4,
	}
	Q.pushQueue(4)
	Q.pushQueue(5)
	Q.pushQueue(3)
	fmt.Println("push 4, 5, 3后为")
	Q.showQueue()

	Q.popQueue()
	Q.popQueue()
	fmt.Println("pop两次后为")
	Q.showQueue()

}