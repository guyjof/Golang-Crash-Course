package main

import "fmt"

func main() {
	purchases := [5]float32{1.90, 5.50, 4.50, 3.20, 2.30}

	mySlice := purchases[:]
	fmt.Println(mySlice) // [1.9 5.5 4.5 3.2 2.3]

	mySlice = append(mySlice, 17.42)
	fmt.Println(mySlice) // [1.9 5.5 4.5 3.2 2.3 17.42]

	myOtherSlice := purchases[1:4]
	fmt.Println(myOtherSlice) // [5.5 4.5 3.2]

	myThirdSlice := purchases[:4] // [1.9 5.5 4.5 3.2]
	// myThirdSlice := purchases[2:] // [4.5 3.2 2.3]
	// fmt.Println(myThirdSlice)

	combine := append(mySlice, myThirdSlice...)
	fmt.Println(combine) // [1.9 5.5 4.5 3.2 2.3 17.42 1.9 5.5 4.5 3.2]

	// deep dive into the slice type
	intArr := [...]int{1, 2, 3}
	fmt.Println(intArr) // [1 2 3]

	// slices are wrappers around arrays that provide more functionality
	// by omitting the size of the array we are creating a slice
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length of intSlice is %d and the capacity is %d\n", len(intSlice), cap(intSlice)) // The length of intSlice is 3 and the capacity is 3
	intSlice = append(intSlice, 7)
	fmt.Printf("The length of intSlice is %d and the capacity is %d\n", len(intSlice), cap(intSlice)) // The length of intSlice is 4 and the capacity is 6
	// this is because the capacity of the slice is doubled when the length exceeds the capacity

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // this is how we can append one slice to another
	fmt.Println(intSlice)                     // [4 5 6 7 8 9]

	var intSlice3 []int32 = make([]int32, 3, 8)                                                          // make is a built-in function that creates a slice of a given type, length, and capacity
	fmt.Printf("The length of intSlice3 is %d and the capacity is %d\n", len(intSlice3), cap(intSlice3)) // The length of intSlice3 is 3 and the capacity is 8
}
