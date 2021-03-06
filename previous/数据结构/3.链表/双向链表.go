package main

import "fmt"

/*
1. 双向链表可以向前或者向后查找。
2. 双向链表可以实现自我删除。
 */

type LanNode struct {
	No int
	Name string
	// next指向下一个节点
	next *LanNode
	// pre指向上一个节点
	pre  *LanNode
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
			// 新节点的pre指向当前临时的tmp节点()
			newNode.pre = tmp
			// 让tmp下一个节点的pre指向newNode
			tmp.next.pre = newNode
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
			// 因为tmp.next.next可能是nil, 那么如果是的话,则nil没有pre, 为nil的时候直接暂停,不需要pre了.
			if tmp.next != nil {
				// 因为tmp.next已经指向了tmp.next.next, 所以这里的tmp.next已经跳过了被删除的节点. 让tmp的下个节点的pre指向tmp
				tmp.next.pre = tmp
			}
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
