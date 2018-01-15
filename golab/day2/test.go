package main

import "fmt"

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(), array values : ", array)
}

// no modify
func testArray() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In testArray(), array values : ", array)
}

func main() {
	str := "hello, 世界."
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	for i, ch := range str {
		fmt.Println(i, ch)
	}
	fmt.Println("vim-go")
	fmt.Println("======================")
	testArray()
	fmt.Println("======================")
	testSlice()
	fmt.Println("======================")
	myslice := make([]int, 5, 10)
	fmt.Println("len(myslice): ", len(myslice))
	fmt.Println("cap(myslice): ", cap(myslice))
	fmt.Println("======================")

	myslice2 := []int{1, 2, 3}
	myslice = append(myslice, myslice2...)
	for _, v := range myslice {
		fmt.Print(v, " ")
	}
	fmt.Println("======================")
	myslice = append(myslice, myslice2...)
	for _, v := range myslice {
		fmt.Print(v, " ")
	}
	fmt.Println("======================")
	myslice = append(myslice, myslice2...)
	myslice = append(myslice, myslice2...)
	for _, v := range myslice {
		fmt.Print(v, " ")
	}
	fmt.Println("len(myslice): ", len(myslice))
	fmt.Println("cap(myslice): ", cap(myslice))
	myslice = append(myslice, myslice2...)
	myslice = append(myslice, myslice2...)
	fmt.Println("len(myslice): ", len(myslice))
	fmt.Println("cap(myslice): ", cap(myslice))
}
