package main

import "fmt"

// Arrays in go are fixed-size sequences of elements of the same type that are indexable and stored in contiguous memory locations
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

	// deep dive into the array type

	var intArray [3]int32      // declare an array of 3 elements of type int32
	fmt.Println(intArray)      // [0 0 0]
	intArray[0] = 1            // [1 0 0]
	fmt.Println(intArray[1:3]) // get a slice of the array from index 1 to 3 (not inclusive) => [0 0]

	// because arrays are stored in contiguous memory locations each element in the array can be accessed in constant time
	// this is because the memory address of the first element is known and the address of any other element can be calculated
	// by adding the size of the element to the address of the first element
	fmt.Println(&intArray[0]) // 0x14000112040
	fmt.Println(&intArray[1]) // 0x14000112044
	fmt.Println(&intArray[2]) // 0x14000112048

	// we can also initialize an array using this syntax
	intArray2 := [3]int32{1, 2, 3} // here we are initializing an array of 3 elements
	// intArray2 := [...]int32{1, 2, 3} // here we are initializing an array of 3 elements but the compiler will count the number of elements to determine the size of the array
	fmt.Println(intArray2) // [1 2 3]
}
