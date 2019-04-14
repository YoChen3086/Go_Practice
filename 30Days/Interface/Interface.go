package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perimeter() float64
}

type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}
func (s square) perimeter() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	measure(s)
	measure(c)
}

// 這個範例很簡單淺顯易懂，
// 首先我們看到 interface 這邊，
// 你會發現他的語法和 struct 很類似，
// 裡面是繼承 area 跟 perimeter 的方法，
// 然後後面利用 struct 建立正方形跟圓形的結構。

// 再往下看，定義的是 area 跟 perimeter 的方法，
// 而引入的參數是前面用 struct 定義的結構，再來是定義 measure 函式，
// 這個就很簡單，只是將當作參數的 interface 繼承的方法印出而已。

// 說明完定義的部份，接下來就很簡單了，
// 主函式這邊我們利用剛剛建立的 struct 建立兩個變數 s, c ，然後呼叫函式。
// 以 measure(s) 來看，
// 因為我們初始化 struct 的數值的時候給予 width: 3, height: 4，
// 所以第一行會印出 {3 4}，再來會印出 12 ，為什麼呢？
// 因為他去呼叫 interface g 繼承的 area() 方法，
// 然後方法是 s.width * s.height 所以自然就會印出 12 囉～
// 因為 3 * 4 = 12 ，是不是很簡單阿。
