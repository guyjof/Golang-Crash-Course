package main

import "fmt"

func main() {
	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	this is the same as the for loop above
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}
	this is the same as the for loop above
	k := 1

	for {
		if k > 5 {
			break
		}
		fmt.Println(k)
		k++
	}
	// best practice is to use the for loop when you know the number of iterations
	// and the while loop when you don't know the number of iterations
}
