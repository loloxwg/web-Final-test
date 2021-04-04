package main

import "sync"

var (
	myRes = make(map[int]int, 20)
	// 声明全局互斥锁
	lock sync.Mutex
)
// 声明一个map，用于存储每个数的阶乘
//var myRes = make(map[int]int, 20)

func main() {
	// 启用20个goroutine同时求1~20以内各个数的阶乘
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
}

// 求n的阶乘，并将结果写进myRes
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()//进行存储操作前，加锁
	myRes[n] = res
	lock.Unlock()  //存储完毕后解锁
}
