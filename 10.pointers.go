package main

import "fmt"

func main() {
	a := 1
	b := a

	b += 1
	fmt.Println(a, b) // 1, 2

	c := 1 // c is an int

	d := &c // d is a pointer to an int

	fmt.Println(c) // 1

	*d += 1

	fmt.Println(c) // 2

	var e int = 1
	var f *int = &e
	fmt.Printf("The value of e is %v and the value of f is %v\n", e, f) // expected output: The value of e is 1 and the value of f is 0xc0000120b8
}
