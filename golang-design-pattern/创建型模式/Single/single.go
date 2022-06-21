package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	// 判断是否为nil
	if singleInstance == nil {
		// 加锁,防止多进程的时候发生竞争
		lock.Lock()
		defer lock.Unlock()
		// 加锁后还得判断下,是否加锁期间singleInstance被赋值
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}
