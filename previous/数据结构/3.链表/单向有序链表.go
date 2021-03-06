package main

import "fmt"

/*
设计一个单向链表，要求插入的数据能够有序。
NO.1 C
No.2 PHP
No.3 Java
No.4 NodeJs

链表会有一个head头，该节点不包含任何数据。
遍历链表，从head开始，取next的，依次取，如果为nil说明到链表尾部。
插入元素，从head开始，取出下一个节点的No进行比较，如果大于则在当前位置插入，让插入的新节点指向下一个节点，临时节点指向新节点。
删除元素，从head开始，取出下一个节点的No进行比较，如果大于则在当前位置删除, 将临时节点的next指向下下个节点。
 */

type LanNode struct {
	No int
	Name string
	// next指向下一个节点
	next *LanNode
}

func newNode(no int, name string) *LanNode{
	return &LanNode{
		No:no,
		Name:name,
		// 默认让next指向下一个节点，所以只需要传入no和string即可。
	}
}

func insertNode(head *LanNode, newNode *LanNode) {
	// 给定一个临时节点，用来遍历
	// 该临时节点为内存引用
	tmp := head
	// 如果当前链表为空，则直接插入newNode
	for {
		if tmp.next == nil {
			tmp.next = newNode
			fmt.Println("已经遍历到链表尾部，插入成功")
			return
		}
		// 当next的No大于新节点的No，表示可以插入
		if tmp.next.No > newNode.No {
			// 让插入的新节点指向next节点
			newNode.next = tmp.next
			// 让tmp的next节点指向新节点
			tmp.next = newNode
			fmt.Println("已经插入到合适位置")
			return
		}

		// 存在相同的情况，则退出
		if tmp.next.No == newNode.No {
			fmt.Println("无法插入该节点，已有相同No")
			return
		}
		// 步进
		tmp = tmp.next
	}
}

func deleteNode(head *LanNode, no int) {
	// 给tmp一个赋值，作为临时节点，用来遍历
	// 这个tmp其实是内存的引用
	tmp := head
	for {
		if tmp.next == nil {
			fmt.Println("不存在待删除的no")
			return
		}
		if tmp.next.No == no {
			// 将tmp当前的节点next指向tmp下一个节点的下一个节点。
			tmp.next = tmp.next.next
			fmt.Println("删除成功")
			return
		}
		// 步进
		tmp = tmp.next
	}

}


// 对节点进行遍历
func listNode(head *LanNode) {
	// 从head的next开始遍历，直到nil为止
	for {
		if head.next == nil {
			fmt.Println("已经遍历到底部")
			return
		} else {
			fmt.Printf("No为%d, Name为%s语言 ==> ",head.next.No, head.next.Name)
		}
		// 步进
		head = head.next
	}
}


func main() {
	// 定义head头，不需要存放任何数据
	headNode := &LanNode{}

	// 定义四个新的节点
	nodeC := newNode(1, "C")
	nodePHP := newNode(2, "PHP")
	nodeJava := newNode(3, "Java")
	nodeNodeJs := newNode(4, "NodeJs")

	// 进行数据插入, 这里乱序插入
	insertNode(headNode, nodeC)
	insertNode(headNode, nodeJava)
	insertNode(headNode, nodeNodeJs)
	insertNode(headNode, nodePHP)

	listNode(headNode)

	// 删除一个数据3
	deleteNode(headNode, 3)
	listNode(headNode)

}
