package main

import (
	"fmt"
	"math"
)

func main() {
	easyFor()
	while()
	endlessFor()

	fmt.Println("If範例")
	fmt.Println(ifExample(2), ifExample(-4))

	fmt.Println("If加敘述")
	fmt.Println(
		ifNarrative(3, 2, 10),
		ifNarrative(3, 3, 20),
	)

	fmt.Println("Switch範例1")
	switchExample(1)
	switchExample(3)
	switchExample(4)

	fmt.Println("Switch範例2")
	switchAnotherExample(6)
}

func easyFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("For迴圈範例")
	fmt.Println(sum)
}

func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("While迴圈範例")
	fmt.Println(sum)
}

func endlessFor() {
	fmt.Println("無窮迴圈範例\n")
	// for {

	// }
}

func ifExample(x float64) string {
	if x < 0 {
		return ifExample(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ifNarrative(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func switchExample(i int) {
	switch i {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4:
		fallthrough
	case 5:
		fmt.Println("5")
	}
}

func switchAnotherExample(i int) {
	// Switch 後的陳述不是必須的，你可以寫 Case 裡面。
	switch {
	case i > 5:
		fmt.Println(i, "> 5")
	case i < 5:
		fmt.Println(i, "< 5")
	case i == 5:
		fmt.Println(i, "= 5")
	}
}
