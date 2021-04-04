package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明了一个带有3个缓存空间的channel
	var channel2 = make(chan int, 3)

	go func() {
		channel2 <- 1
		channel2 <- 2
		channel2 <- 3
		fmt.Println("我发送了3个数据")
		// 当发送第四个值的时候，goroutine阻塞
		channel2 <- 4
		fmt.Println("我发送了第4个数据")
	}()
	time.Sleep(time.Second)
}