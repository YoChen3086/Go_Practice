package main

import "fmt"

// 錯誤處理
// 以前的 try...catch...finally 功能雖然好用，
// 但是會導致整個結構變得很糟糕，
// 維護又很不方便，
// 而 Go 改良的這個問題。

// 首先要處理錯誤就必須要加上 defer 這樣才能正確的抓到，
// 再這個例子我們藉由 panic 產生的個錯誤並且值為 1 ，
// 這裡要注意的是 Go 的 panic 的錯誤真的是錯誤，
// 程式一定會發生錯誤然後關閉，
// 除非你有使用 recover 會回覆他的數值。

func main() {
	defer func() {
		fmt.Println("first")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end")
	}()
	f()
}

func f() {
	fmt.Println("test")
	panic(1)
	fmt.Println("test2")
}
