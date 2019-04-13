package main

import (
	"fmt"
	"time"
)

// Go 很酷的特色 Goroutine ，他類似於其他語言的 Thread。
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	time.Sleep(time.Second * 1) // 暫停一秒鐘
}
