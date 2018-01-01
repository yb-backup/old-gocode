package main

import (
	"fmt"
)

func count(ch chan int) {
	fmt.Println("counting")
	ch <- 1
}

func TestChannel() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
}
