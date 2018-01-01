package main

import "fmt"

type Rect struct {
	x, y float64
	w, h float64
	area float64
}

func (r *Rect) Area() float64 { // 可以修改rect对象的值
	r.area = r.w * r.h
	return r.area
}

func (r Rect) Area2() float64 { // 不能修改rect对象的值
	r.area = r.w * r.h
	return r.area
}

func main() {
	rect := &Rect{w: 100, h: 50}
	fmt.Println(rect.Area())
	fmt.Println(rect.area)
	rect2 := &Rect{w: 100, h: 50}
	fmt.Println(rect2.Area2())
	fmt.Println(rect2.area)
}
