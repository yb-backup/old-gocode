package main

func test1() {
	var v1 int
	var v2 string
	var v3 [10]int // 数组
	var v4 []int   //切片
	var v5 struct {
		f int
	}
	var v6 *int // 指针
	var v7 map[string]int
	var v8 func(a int) int
	var (
		v9  int
		v10 string
	)
}

func test2() {
	var v1 int = 10
	var v2 = 10
	v3 := 10
	// v2 := 20 // err
	var v4 int
	// v4 = "s" // err
	v4 = 10
}

func test3() {
	const Pi float64 = 3.1415926
	const zero = 0.0
	const (
		size int64 = 1024
		eof        = -1
	)
	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "foo"
}

func test4() {
	var v1 bool
	v1 = true
	v2 := (1 == 2)
}
