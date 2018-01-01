package main

import (
	"fmt"
	"time"
)

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func TestAdd() {
	for i := 0; i < 10; i++ {
		go add(i, i)
	}
	time.Sleep(3 * 100 * time.Microsecond)
}
