package main

import "fmt"

func main() {
	arrayExample1()
	arrayExample2()
}

func arrayExample1() {
	fmt.Println("\nArray範例一")

	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func arrayExample2() {
	fmt.Println("\nArray範例二")

	// 以下兩種初始化陣列的方式相同
	// var x [4]float64
	// x[0] = 23
	// x[1] = 45
	// x[2] = 33
	// x[3] = 21

	x := [4]float64{
		23,
		45,
		33,
		21, // 逗號必須保留
	}

	var total float64 = 0
	// 以下這三種走遍Array寫法都可，但range是最簡潔的
	// for i := 0; i < 4; i++ {
	// 	total += x[i]
	// }

	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }

	for _, value := range x {
		total += value
	}

	fmt.Println(x)
	// total是float，len(x)是int，必須強制轉型
	fmt.Printf("Average = %g\n", total/float64(len(x)))
}
