package main

import "fmt"

func breakExample() {
	/* 定义局部变量 */
	var a int = 10

	fmt.Println("break")

	/* for 循环 */
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}
}

func continueExample() {
	/* 定义局部变量 */
	var a int = 10

	fmt.Println("continue")

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

func gotoExample() {
	/* 定义局部变量 */
	var a int = 10

	fmt.Println("loop")

	/* 循环 */
Loop:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto Loop
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

func main() {
	breakExample()
	continueExample()
	gotoExample()
}
