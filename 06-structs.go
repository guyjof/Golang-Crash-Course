package main

import "fmt"

func main() {
	type Animal struct {
		class  string
		age    int
		gender string
	}
	// order of fields doesn't matter
	teddy := Animal{
		age:    5,
		class:  "mammal",
		gender: "male",
	}

	teddy.age += 1

	fmt.Println(teddy)
	// order of fields matters when not using field names
	leo := Animal{"mammal", 3, "male"}

	fmt.Println(leo)
}
