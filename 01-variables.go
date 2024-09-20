package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// use Println to print to the console in a new line
	fmt.Println("Hello, World!")

	var favoriteBook = "The Count of Monte Cristo"
	fmt.Println(favoriteBook) // expected output: The Count of Monte Cristo
	favoriteBook = "Harry Potter"
	fmt.Println(favoriteBook) // expected output: Harry Potter

	// favoriteBook = 1 // this line throws an error => cannot use 12 (untyped int constant) as string value in assignment

	var isBook bool                             // default value is false
	var numPages int                            // default value is 0
	fmt.Println(favoriteBook, isBook, numPages) // expected output: Harry Potter false 0

	// shorthand variable declaration AKA compound creation
	var name, age, isStudent = "John Doe", 25, true
	fmt.Println(name, age, isStudent) // expected output: John Doe 25 true

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

	// Deep dive into the int type
	var a int = 10 // int is a signed integer type that has a size of 32 or 64 bits depending on the underlying platform

	var b int8 = 127  // int8 is a signed integer type that has a size of 8 bits min value of -128 and max value of 127
	var c uint8 = 255 // uint8 is an unsigned integer type that has a size of 8 bits min value of 0 and max value of 255

	var (
		d int16 = 32767               // int16 is a signed integer type that has a size of 16 bits min value of -32768 and max value of 32767
		e int32 = 2147483647          // int32 is a signed integer type that has a size of 32 bits min value of -2147483648 and max value of 2147483647
		f int64 = 9223372036854775807 // int64 is a signed integer type that has a size of 64 bits min value of -9223372036854775808 and max value of 9223372036854775807
	)

	fmt.Println(a, b, c, d, e, f) // expected output: 10 127 255 32767 2147483647 9223372036854775807

	// Deep dive into the float type

	var g float32 = 12345678.9 // this will be rounded to 12345679 because float32 has a precision of 7 decimal places
	var h float64 = 12345678.9 // float64 has a precision of 15 decimal places so it will be 12345678.9

	fmt.Println(g, h) // expected output: 12345679.0 12345678.9

	var float32Num float32 = 10.1
	var int32Num int32 = 10
	// var result float32 = float32Num + int32Num // this line throws an error, cannot use int32Num (type int32) as type float32 in arithmetic operation
	var result float32 = float32Num + float32(int32Num)
	fmt.Println(result) // expected output: 20.1

	// Deep dive into the string type
	var message string = "Hello"
	fmt.Println(len(message)) // expected output: 5
	fmt.Println(message[0])   // expected output: 72 => H in ASCII
	fmt.Println(len("Ý"))     // expected output: 2 => Ý in ASCII is stored in 2 bytes
	// The built-in Len function returns the number of bytes in a string and not the number of characters
	// go uses utf-8 encoding characters outside the basic ASCII character set are stored with more than a single byte so taking
	// the length of this "A" ASCII character will give you a one but taking the length of "ý" for example gives you 2
	// to get the number of characters in a string you can use the utf8 package from the standard library to get the number of runes in a string
	fmt.Println(utf8.RuneCountInString("Ý")) // expected output: 1

}
