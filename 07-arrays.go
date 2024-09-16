package main

import "fmt"

func main() {
	purchases := [5]float32{1.90, 5.50, 4.50, 3.20, 2.30}

	fmt.Println(purchases[0]) // 1.9

	purchases[0] = 2.50

	fmt.Println(purchases) // [2.5 5.5 4.5 3.2 2.3]

	// var sales [4]float32 // [0 0 0 0]
	// sales[0] = 1.50
	// fmt.Println(sales) // [1.5 0 0 0]

	sales := [...]float32{12.34, 56.78, 91.011} // [...] means the compiler will count the number of elements to determine the size of the array
	fmt.Println(sales)                          // [12.34 56.78 91.011]

	for i := 0; i < len(sales); i++ {
		fmt.Println(sales[i])
	}
}
