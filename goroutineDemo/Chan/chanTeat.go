package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan int)
	var send = 6666
	var receive int
	// 关键字go后跟的是需要被并发执行的代码块，它由一个匿名函数代表
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以
	go func() {
		channel <- send
		fmt.Println("数据已发送")
	}()
	// 这里让主线程休眠1秒钟
	// 以使上面的goroutine有机会被执行

	receive=<-channel
	time.Sleep(1 * time.Second)
	fmt.Println(receive)
}
