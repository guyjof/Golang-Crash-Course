package main

import "fmt"

func main() {
	// use Println to print to the console in a new line
	fmt.Println("Hello, World!")

	// var favoriteBook = "The Count of Monte Cristo"
	// fmt.Println(favoriteBook) // expected output: The Count of Monte Cristo
	// favoriteBook = "Harry Potter"
	// fmt.Println(favoriteBook) // expected output: Harry Potter

	// favoriteBook = 1 // this line throws an error => cannot use 12 (untyped int constant) as string value in assignment

	// var isBook bool                             // default value is false
	// var numPages int                            // default value is 0
	// fmt.Println(favoriteBook, isBook, numPages) // expected output: Harry Potter false 0

	// shorthand variable declaration AKA compound creation
	// var name, age, isStudent = "John Doe", 25, true

	// block creation
	// var (
	// 	name      = "John Doe"
	// 	age       = 25
	// 	isStudent = true
	// )

	// short variable declaration
	// name := "John Doe"

	// short variable declaration with multiple variables
	// name, age, isStudent := "John Doe", 25, true
	// name := "Jane Doe" // this line throws an error, name has already been declared
	// gender, name := "male", "Jane Doe" // this will work because I'm declaring a new variable: gender

	const pi = 3.14159
	fmt.Printf("The value of pi is %v and the type is %T\n", pi, pi) // expected output: The value of pi is 3.14159 and the type is float64
	// pi = 2.71828 // this line throws an error, cannot assign to pi because it is a constant
}
