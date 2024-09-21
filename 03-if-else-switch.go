package main

import (
	"fmt"
	"math"
)

func main() {
	age := 25
	animal := "dog"

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age == 17 {
		fmt.Println("You are almost an adult")
	} else {
		fmt.Println("You are a minor")
	}

	switch animal {
	case "dog":
		fmt.Println("You have a dog")
	case "cat":
		fmt.Println("You have a cat")
	default:
		fmt.Println("You have an animal")
	}

	// short if statement
	if x := math.Pow(2, 3); x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}
	// fmt.Println(x) // x is not available here because it is declared in the if statement
}
