package main

import "time"
import "fmt"

// Channel 有一個類似 Switch 的流程控制「Select」，它只能應用於 Channel 讓我們一起來看看。
// <strong>Select</strong>

// Go by Example 有一個很有意思的範例，我們一起來看看

// 這邊用 go 建立兩個 goroutine 分別將 one 和 two 傳給 c1, c2，
// 下面主函式的 for 回會將 1, 2 透過 select 流程控制來接收 channel 的訊息再印出。

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
