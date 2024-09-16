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
}
