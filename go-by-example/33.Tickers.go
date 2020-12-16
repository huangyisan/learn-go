package main

import (
	"fmt"
	"time"
)

// NewTicker 将会还不停进行间隔执行
func main() {
	ticker1 := time.NewTicker(time.Millisecond * 500)

	go func() {
		// 间隔500毫秒执行，这边可以看做每过500ms往C塞入一个数据
		for t := range ticker1.C {
			fmt.Println("Tick at ", t)
		}
	}()
	time.Sleep(time.Millisecond * 16000)
	ticker1.Stop()
	fmt.Println("ticker1 stop")
}