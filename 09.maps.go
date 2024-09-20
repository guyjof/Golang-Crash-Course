package main

import "fmt"

func main() {
	myCart := make(map[string]float32)

	myCart["apple"] = 1.99
	myCart["banana"] = 0.99
	myCart["orange"] += 1.49

	fmt.Println(myCart) // map[apple:1.99 banana:0.99 orange:1.49]

	// Reading a value from a map
	applePrice, ok := myCart["apple"]
	fmt.Println(applePrice, ok) // 1.99, true
	applePrice, ok = myCart["book"]
	fmt.Println(applePrice, ok) // 0, false

	// Deleting a value from a map
	delete(myCart, "apple")
	fmt.Println(myCart) // map[banana:0.99 orange:1.49]

	// deep dive into the map type
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // map[]

	var myMap2 = map[string]uint8{"one": 1, "two": 2, "three": 3}

	for key, value := range myMap2 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
