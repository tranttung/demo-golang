package main

import "fmt"

func main() {

	a := 1000
	var p *int
	p = &a
	fmt.Println(p)
	fmt.Printf("%T", p)
	b := 100.23
	p2 := new(float64)
	p2 = &b
	fmt.Println()
	fmt.Println(p2)
	fmt.Printf("%T", p2)
	fmt.Println()
	fmt.Println("=========")
	var pointer *int
	x := 100
	pointer = &x
	fmt.Println(a)
	*pointer = 999 // x := 999
	fmt.Println(pointer)
	fmt.Println(x)
	fmt.Println()

	// pointer -> array
	array := [3]int{1, 2, 3}
	var pointer2 *[3]int
	pointer2 = &array
	fmt.Println(pointer2)
}
