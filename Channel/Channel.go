package main

import "fmt"

// 這邊簡單建立了一個 message 的 channel ，
// 可以傳輸字串，然後用 go 來 call goroutine 執行函式，
// 然後 msg 負責接收 messages 的傳輸資料，
// goroutine 執行的函式裡面傳 "ping" 到 messages 這個 channel 裡面，
// 再由 message 傳給 msg 變數印出。
// 很簡單直覺對吧！透過這個方法就可以簡單的讓 Goroutine 可以溝通！
func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
