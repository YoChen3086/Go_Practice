package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
	appendExample()
	copyExample()
	specialExample()
}

func example1() {
	var slice []float64
	fmt.Printf("範例一: %g\n", slice)
}

func example2() {
	slice := make([]float64, 5)
	fmt.Printf("範例二: %g\n", slice)
}

func example3() {
	slice := make([]float64, 5, 10)
	fmt.Printf("範例三: %g\n", slice)
}

func example4() {
	arr := [5]float64{1, 2, 3, 4, 5}
	// 會參考其他Slice或陣列
	slice := arr[0:5]
	fmt.Printf("範例四: %g\n", slice)
}

func appendExample() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Append範例")
	fmt.Println(slice1, slice2)
}

func specialExample() {
	fmt.Println("特殊範例")
	arr := []float64{1, 2, 3, 4, 5}
	arr2 := arr[2:4]

	fmt.Println("arr", arr)
	fmt.Println("arr2", arr2)
	arr = append(arr, 6)
	fmt.Println("\nChange Origin Data\n")
	fmt.Println("arr", arr)
	fmt.Println("arr2", arr2)

	fmt.Println("\nAnother Case")

	arr3 := []float64{1, 2, 3, 4, 5}
	arr4 := append(arr3, 6)
	fmt.Println("\nChange Origin Data\n")
	arr3 = append(arr3, 7)

	fmt.Println("arr3", arr3)
	fmt.Println("arr4", arr4)
}

func copyExample() {
	// slice 內建的另一個好用的方法 copy，
	// 可以將一個 slice 的值複製到另外一個 slice，
	// 但是無法超過容量限制，
	// 下面的範例將容量為 3 的 slice 的數值 1, 2, 3 複製到 slice2 ，
	// 但是 slice 2 的容量只有 2，
	// 所以實際上 slice2 的值是 1, 2。
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println("Copy範例")
	fmt.Println(slice1, slice2)
}
