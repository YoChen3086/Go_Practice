package main

import "fmt"

// func zero(x int) {
// 	x = 0
// }
// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // x is still 5
// }

// 這樣子就可以用了耶，要注意的是 zero(&x)，這邊有一個 「 & 」，
// 這個很重要不可以忘記，如果沒有這個就找不到指標囉，這個再讀取的時候一定會用到。

// 另外我再說詳細一點給沒碰過的人看，
// 所謂的指標就是指到記憶體位置，
// 上面的例子函式內不是宣告了一個指標 xPtr 嗎？
// 以這邊來看 x 帶進來後就會去看他的「記憶體位置」，
// 然後我們再下面的 *xPtr = 0 更改了這個「記憶體位置」的「值」，
// 所以回到外面印出的時候就已經是 0 了，
// 因為記憶體位置上的值已經被更改囉！

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}
