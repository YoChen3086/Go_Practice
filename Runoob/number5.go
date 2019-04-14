package main

import "fmt"

func ifExample() {
	/* 定义局部变量 */
	var a int = 10

	fmt.Println("if語句")

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func ifElseExample() {
	/* 局部变量定义 */
	var a int = 100

	fmt.Println("if...else語句")

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func depthIfExample() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Println("多層if語句")

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}

func main() {
	ifExample()
	ifElseExample()
	depthIfExample()
}
