package main

import "fmt"

func main() {
	// Arithmetic operators: +, -, *, /, %
	num1, num2 := 10, 5

	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quot := num1 / num2
	remainder := num1 % num2

	fmt.Println(sum, diff, product, quot, remainder) // expected output: 15 5 50 2 0

	num1, num2 = 10, 3
	quot = num1 / num2
	fmt.Println(quot) // expected output: 3, because both num1 and num2 are integers and the result is also an integer (3.3333 -> 3)

	var num3 float32 = 10
	var num4 float32 = 3
	intQuot := num3 / num4 // expected output: 3.33
	fmt.Println(intQuot)

	// Relational operators: ==, !=, >, <, >=, <=
	fmt.Println(num1 == num2) // expected output: false
	fmt.Println(num1 != num2) // expected output: true
	fmt.Println(num1 > num2)  // expected output: true
	fmt.Println(num1 < num2)  // expected output: false
	fmt.Println(num1 >= num2) // expected output: true
	fmt.Println(num1 <= num2) // expected output: false

	// Logical operators: &&, ||, !
	name := "Guy"
	age := 25

	inviteToParty := name == "Guy" && age >= 18 // only invite to party if the name is Guy and the age is 18 or older
	fmt.Println(inviteToParty)                  // expected output: true
	// invite if the name is Guy or the age is 18 or older
	fmt.Println(name == "Guy" || age >= 18) // expected output: true
	// invite if the name is not Guy
	fmt.Println(name != "Guy") // expected output: false
}
