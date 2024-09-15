package main

import "fmt"

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
}
