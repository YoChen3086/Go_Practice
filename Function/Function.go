package main

import "fmt"

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

// 多數值的返回
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("world", "hello")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
