package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
// 操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

// Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：
// 如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭。
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	// 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)

	t := make(chan int, 10)
	go fibonacci(cap(t), t)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range t {
		fmt.Println(i)
	}
}
