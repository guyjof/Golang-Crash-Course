package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 11
	var denominator int = 2

	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		return
	} else if remainder == 0 {
		fmt.Printf("The result of %d divided by %d is %d\n", numerator, denominator, result)
	} else {
		fmt.Printf("The result of %d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		err := errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
