package main

import "fmt"

func main() {
	myCart := make(map[string]float32)

	myCart["apple"] = 1.99
	myCart["banana"] = 0.99
	myCart["orange"] += 1.49

	fmt.Println(myCart) // map[apple:1.99 banana:0.99 orange:1.49]

	// Reading a value from a map
	// applePrice, found := myCart["apple"]
	// fmt.Println(applePrice, found) // 1.99, true
	applePrice, found := myCart["book"]
	fmt.Println(applePrice, found) // 0, false

	// Deleting a value from a map
	delete(myCart, "apple")
	fmt.Println(myCart) // map[banana:0.99 orange:1.49]
}
