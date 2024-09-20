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

	// deep dive into pointers

	// var p *int32                                      // initial value of p is <nil>
	var p *int32 = new(int32) // initial value of p is 0
	var i int32
	fmt.Printf("The value p is the address of: %v\n", p) // expected output: 0x1400000e0f8
	fmt.Printf("The value p points to is: %v\n", *p)     // expected output: 0
	// now if we want to change the value p points to
	*p = 42
	fmt.Printf("The value p points to now is: %v\n", *p) // expected output: 42
	// we can also set the address of p to the memory address of i
	// by using the & operator to get the address of i
	p = &i
	fmt.Printf("The value p is the address of: %v\n", p) // expected output: 0x1400010402c
	fmt.Printf("The value p points to is: %v\n", *p)     // expected output: 0
}
