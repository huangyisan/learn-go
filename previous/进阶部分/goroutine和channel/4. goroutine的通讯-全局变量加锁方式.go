package main

import (
	"fmt"
	"sync"
	"time"
)

/*
不同goroutine之间如何通讯
1. 全局变量加锁同步
2. 使用channel
 */

 // 使用全局变量加锁同步
var (
	myMap = make(map[int]int64, 10)

	// lock 是一个全局的互斥锁
	lock sync.Mutex
)


// 计算n的阶乘，将结果放入myMap中

func test(n int) {
	res := 1
	for i :=1; i<=n; i++ {
		res *= i
	}

	// 将res结果放入myMap
	// 这个动作执行之前加锁
	lock.Lock()
	myMap[n] = int64(res)
	// 执行完后解锁
	lock.Unlock()
}

func main() {
	//启动多个协程
	for i := 1; i<=20; i++ {
		go test(i)
	}

	// 防止主进程先结束，手工10s等待。
	time.Sleep(time.Second * 10)

	// 从程序设计上能知道10s肯定能结束所有协程，但主线程并不知道，因此底层还是可能会产生资源争夺，因此还是需要加入互斥锁来解决。
	// 等test方法释放锁之后, 下面获取锁然后执行.
	lock.Lock()
	for k,v := range myMap {
		fmt.Printf("key is %v, value is %v\n",k,v)
	}
	lock.Unlock()

}