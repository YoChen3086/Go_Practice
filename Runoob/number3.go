package main

import (
	"fmt"
	"unsafe"
)

// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
const (
	LENGTH  int = 10
	WIDTH   int = 5
	a, b, c     = 1, false, "str"
)

// 常量可以用len(), cap(),
// unsafe.Sizeof()函数计算表达式的值。
// 常量表达式中，函数必须是内置函数，否则编译不过：
const (
	d = "def"
	e = len(d)
	f = unsafe.Sizeof(d)
)

// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
// iota 可以被用作枚举值：
func iotaExample1() {
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)
}

// iota 表示从 0 开始自动加 1，
// 所以 i=1<<0, j=3<<1（<< 表示左移的意思），
// 即：i=1, j=6，这没问题，
// 关键在 k 和 l，
// 从输出结果看 k=3<<2，l=3<<3。

// 简单表述:

// i=1：左移 0 位,不变仍为 1;
// j=3：左移 1 位,变为二进制 110, 即 6;
// k=3：左移 2 位,变为二进制 1100, 即 12;
// l=3：左移 3 位,变为二进制 11000,即 24。
func iotaExample2() {
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println(i, j, k, l)
}

func main() {
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

	println(d, e, f)

	iotaExample1()
	iotaExample2()
}
