package main

import "fmt"

func first() {
	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func second() {
	// 第二种，根据值自行判定变量类型。

	// 声明一个变量并初始化
	var a = "RUNOOB"

	// 没有初始化就为零值
	var b int

	// bool 零值为 false
	var c bool

	fmt.Println(a, b, c)
}

func third() {
	// 第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：

	var intVal int

	// 这时候会产生编译错误
	// intVal := 1

	// 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	intVal, intVal1 := 1, 2

	fmt.Println(intVal, intVal1)
}

func multi() {
	//类型相同多个变量, 非全局变量
	var vname1, vname2, vname3 int

	// 和 python 很像,不需要显示声明类型，自动推断
	var vname4, vname5, vname6 = "123", 1, false

	// 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
	vname7, vname8, vname9 := "234", 2, true

	// 这种因式分解关键字的写法一般用于声明全局变量
	var (
		vname10 string
		vname11 int
	)

	fmt.Println(vname1, vname2, vname3)
	fmt.Println(vname4, vname5, vname6)
	fmt.Println(vname7, vname8, vname9)
	fmt.Println(vname10, vname11)
}

func main() {
	first()
	second()
	third()
	multi()
}
