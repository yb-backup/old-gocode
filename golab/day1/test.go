package main

import "fmt"

// func variables() {
// 	var v1 int
// 	var v2 string
// 	var v3 [10]int
// 	type v5 struct {
// 		f int
// 		g string
// 	}
// 	var v6 *int
// 	var v7 map[string]int
// 	var v8 map[int]int
// 	var v9 func(a int) int
// 	var (
// 		v10 int
// 		v11 string
// 	)
// }

func f2() {
	var v1 int = 10
	var v2 string = "10"
	v3 := 10
	v4 := "10"
	// fmt.Println("%s - %s - %s - %s", v1, v2, v3, v4)
	fmt.Printf("%d - %s - %d - %s", v1, v2, v3, v4)
}

func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

func f3() {
	var v1 int = 10
	// v1 := 10 # Error
	_, _, nickName := GetName()
	// fmt.Println("%s - %s", v1, nickName)
	fmt.Println(nickName)
	fmt.Printf("%d - %s", v1, nickName)
}

func main() {
	f2()
	f3()
	fmt.Println("vim-go")
}
