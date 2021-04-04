package main

import (
	"fmt"
	"time"
)

func a() {
	go repeat("gugugu", 100)

	repeat("阿巴阿巴", 5)
}

func repeat(str string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond)
	}
}
