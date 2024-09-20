package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Generics in Go are a way to write code that can work with any type of data.

func main() {
	// problem()
	// solution()
	unmarshalJSON()
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount uint8
}

func unmarshalJSON() {
	// Load the JSON data into a slice of contactInfo and purchaseInfo.
	var contacts []contactInfo = LoadJSON[contactInfo]("./json/contactInfo.json")
	fmt.Println(contacts)

	var purchases []purchaseInfo = LoadJSON[purchaseInfo]("./json/purchaseInfo.json")
	fmt.Println(purchases)
}

func LoadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath) // Read the JSON file.
	loaded := []T{}                  // Create a slice of the type T.
	json.Unmarshal(data, &loaded)    // Unmarshal the JSON data into the slice.
	return loaded                    // Return the slice.
}

// The problem:
func problem() {
	// We have a function that sums the elements of a slice of integers.
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice)) // 6
}

func sumIntSlice(slice []int) int {
	var sum int
	for i := range slice {
		sum += slice[i]
	}
	return sum
}

func solution() {
	// We can use generics to create a function that sums the elements of a slice of any type of data.
	var intSlice = []int{1, 2, 3}
	sumSlice(intSlice) // 6

	var floatSlice = []float32{1.1, 2.2, 3.3}
	sumSlice(floatSlice) // 6.6000004
}

func sumSlice[T int | float32 | float64](slice []T) {
	var sum T
	for _, v := range slice {
		sum += v
	}
	fmt.Println(sum)
}
