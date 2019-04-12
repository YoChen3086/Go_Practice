package main

import "fmt"

func main() {
	// var x map[string]int

	mapStringExample()
	mapIntExample()
	deleteExample()
	example()
	checkEmpty()
	simpleStructure()
	complexExample()
}

func mapStringExample() {
	x := make(map[string]int)
	x["key"] = 10
	x["key2"] = 20
	fmt.Println(x)
}

func mapIntExample() {
	x := make(map[int]int)
	x[1] = 10
	x[2] = 20
	fmt.Println(x)
}

func deleteExample() {
	x := make(map[int]int)
	x[1] = 10
	x[2] = 20
	delete(x, 1)
	fmt.Println(x)
	// string format用法
	// https://golang.org/pkg/fmt/
	fmt.Printf("長度: %d\n", len(x))
}

func example() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	fmt.Println(elements["Al"])
}

func checkEmpty() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	_, ok := elements["Ai"]
	fmt.Printf("Elements['Ai']存不存在 : %t\n", ok)

	_, ok = elements["Ne"]
	fmt.Printf("Elements['Ne']存不存在 : %t\n", ok)
}

func simpleStructure() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)
}

func complexExample() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	fmt.Println(elements)
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
