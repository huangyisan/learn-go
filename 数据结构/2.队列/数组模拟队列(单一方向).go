package main

import "fmt"

/*

 */
type Queue struct {
	head int
	tail int
	maxCap int
	array [4]int
}

func (q *Queue) isFull() bool {
	// q.tail为下标，队列最大容量需要 - 1
	if q.tail == q.maxCap - 1 {
		fmt.Println("队列已满")
		return false
	}
	return true
}

func (q *Queue) isEmpty() bool {
	if q.head == q.tail {
		fmt.Println("队列为空")
		return false
	}
	return true
}


func (q *Queue) pushQueue(value int) {
	// 先判断队列是否满 true为不满
	if q.isFull() {
		q.tail += 1
		q.array[q.tail] = value
	}
}

func (q *Queue) popQueue() {
	// 先判断队列是否为空 true为不空
	if q.isEmpty() {
		q.head += 1
		fmt.Println("弹出元素为")
		fmt.Println(q.array[q.head])
	}
}

func (q *Queue) showQueue() {
	// 获取从head+1到tail+1位切片, 所以需要从head+1号位进行遍历到tail的位置。
	// 需要判断是否为空
	if q.isEmpty() {
		// 因为被head指向的元素是已经被视为弹出了,所以要从head的下一位开始算.
		tmpHead := q.head + 1
		for i:=tmpHead; i<=q.tail;i++ {
			fmt.Println(q.array[i])
		}
	} else {
		fmt.Println("空队列")
	}
}

func main() {
	Q := &Queue{
		head:-1,
		tail:-1,
		maxCap:4,
	}
	Q.pushQueue(1)
	Q.pushQueue(2)
	Q.pushQueue(3)
	Q.pushQueue(3)
	Q.pushQueue(4)
	Q.popQueue()
	Q.showQueue()
}