package main

import "fmt"

func main() {
	var myArray = [5]int{1, 2, 3, 4, 5}
	var mySlice = myArray[:3]
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println()
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
	for _, v := range myArray[3:] {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// oldSlice := []int{1, 2, 3, 4, 5}
	// fmt.Println("oldSlice cap: ", cap(oldSlice))
	// newSlice := oldSlice[:7] // err, cap(oldSlice) == 5
	// for _, v := range newSlice {
	// 	fmt.Print(v, " ")
	// }
	// fmt.Println()
	var oldSlice = make([]int, 5, 10)
	oldSlice = append(oldSlice, 1, 2) // 返回一个新的切片
	fmt.Println("oldSlice cap: ", cap(oldSlice))
	newSlice := oldSlice[:7]
	for _, v := range newSlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
