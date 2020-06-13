package main

import (
	"fmt"
	"errors"
)

// 声明队列结构体, 四个要素, 长度, head, tail, 数组
type queue struct {
	head int
	tail int
	array [4]int
	maxCap int
}


func (q *queue) isFull() bool {
	// 判断队列是否满
	if (q.tail + 1) % q.maxCap == q.head {
		fmt.Println("队列已满")
		return true
	}
	return false
}

func (q *queue) isEmpty() bool {
	// 判断队伍是否为空
	if q.tail == q.head {
		fmt.Println("队列为空")
		return true
	}
	return false
}

func (q *queue) pushQueue(value int)  {

	if q.isFull() {
		fmt.Println("队列满了")
		//// 此时队列满了, 需要取模后重新开始
		//q.tail = (q.tail + 1) % q.maxCap
		//q.array[q.tail] = value
		//q.tail += 1
	} else {
		q.array[q.tail] = value
		// 自增因为存在循环情况, 所以要取模才不会溢出
		q.tail = (q.tail + 1) % q.maxCap
	}


}

func (q *queue) popQueue() error{
	// 弹出之前先判断数组是否为空
	if q.isEmpty() {
		fmt.Println("队列为空")
		return errors.New("队列为空")
	}

	val := q.array[q.head]
	// 自增因为存在循环情况, 所以要取模才不会溢出
	q.head = (q.head + 1) % q.maxCap

	fmt.Printf("弹出了%v\n", val)
	return nil

}

func (q *queue) showQueue() {
	// 因为可能tail会进行下一轮,导致tail的下标比head小, 所以要先得到队列中当前有多少个元素, 那么以head为起点, 循环这些元素个数次数, 则就是队列总的元素内容了
	count := (q.tail + q.maxCap - q.head) % q.maxCap
	if count == 0 {
		fmt.Println("当前为空队列")
		return
	}
	tempHead := q.head
	// 循环元素个数的次数, 因为i的起始值为0, 如果count是5, 那么其实是0,1,2,3,4 所以是i<count,而非i<=count
	for i:=0 ; i<count; i++ {
		fmt.Printf("q.array[%v], 值为%v\n", tempHead, q.array[tempHead])
		// 然后tempHead +1, 但因为是循环, 到了最大空间后, 又从第一个元素开始, 所以要取模
		tempHead = (tempHead + 1) % q.maxCap
	}
}

func main() {
	Q := &queue{
		head:0,
		tail:0,
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