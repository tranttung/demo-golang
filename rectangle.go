package main

import (
	"fmt"
)

type Square struct {
	X float64
}

// Method with receiver `Square`
func (s Square) Acreage() float64 {
	return s.X * s.X
}

type Rectangle struct {
	X, Y float64
}

// Method with receiver `Rectangle`
func (r Rectangle) Acreage() float64 {
	return r.X * r.Y
}

func main() {
	s := Square{4}
	r := Rectangle{3, 4}

	fmt.Println("Diện tích hình vuông là : ", s.Acreage())
	fmt.Println("Diện tích hình chữ nhật là : ", r.Acreage())
}
