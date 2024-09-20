package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string = "résumé"
	fmt.Println(myString) // résumé
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed) // expected output: 114, uint8
	fmt.Println("Index | Value")
	fmt.Println("------|------")
	for i, v := range myString {
		fmt.Printf("  %d   |   %c  \n", i, v)
	}
	// expected output:
	// Index | Value
	// ------|------
	//   0   |   r
	//   1   |   é
	//   3   |   s
	//   4   |   u
	//   5   |   m
	//   6   |   é
	// because é is a rune and takes up two bytes
	// so the index of é takes up two bytes, 1 and 2
	// so the index of s is 3 and the index of u is 4

	var myRunes []rune = []rune(myString)
	fmt.Println(myRunes) // [114 233 115 117 109 233]
	var rune rune = 'a'
	fmt.Printf("%v, %T\n", rune, rune) // expected output: 97, int32

	//strings are immutable
	// myString[0] = 'R' // this line throws an error, cannot assign to myString[0]

	var strSlice = []string{"a", "b", "c"}
	var catStr = ""
	// using this method is not efficient
	// because strings are immutable
	// and every time we concatenate a string
	// a new string is created in memory
	// and the old string is garbage collected
	for _, v := range strSlice {
		catStr += v
	}
	fmt.Println(catStr) // abc

	// using this method is more efficient
	// because we are using the strings.Builder type
	// which is more efficient for string concatenation
	var strBuilder strings.Builder
	for _, v := range strSlice {
		strBuilder.WriteString(v)
	}
	catStr = strBuilder.String()
	fmt.Println(catStr) // abc
}
