package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now() // 返回当前的时间戳
	ch := make(chan bool,1) // 这里make了一个无缓存的channel
	go func ()  {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true  // 无缓冲，发送方阻塞直到接收方接收到数据。
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) //计算从st声明到现在的时间
	time.Sleep(time.Second * 5)
}
