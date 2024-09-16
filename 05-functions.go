package main

import "fmt"

func countDown(num int) {
	for i := num; i >= 0; i-- {
		fmt.Println(i)
	}
}

func sum(a, b int) int {
	return a + b
}

func main() {
	countDown(5)
	sumResult := sum(5, 5)
	fmt.Println(sumResult)
}
