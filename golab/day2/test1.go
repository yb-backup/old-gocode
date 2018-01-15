package main

import (
	"fmt"
)

func testSlice() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var myslice []int = myArray[:5]
	var myslice1 []int = myArray[:]
	var myslice2 []int = myArray[1:5]
	var myslice3 []int = myArray[5:]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of myslice: ")
	for _, v := range myslice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	fmt.Println("\nElements of myslice1: ")
	for _, v := range myslice1 {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of myslice2: ")
	for _, v := range myslice2 {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of myslice3: ")
	for _, v := range myslice3 {
		fmt.Print(v, " ")
	}
}
