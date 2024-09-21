package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func readFile(filename string) error {
	file, err := os.Open(filename) // Open the file.
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}

	defer file.Close() // Close the file. This will be executed after the function returns.

	data, err := io.ReadAll(file) // Read the file using readAll
	if err != nil {
		return fmt.Errorf("Error reading file: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

func processData(data []uint8) {
	start := time.Now()
	defer func() {
		fmt.Println("Time taken:", time.Since(start))
	}()

	for _, v := range data {
		fmt.Println(v)
		time.Sleep(time.Millisecond * 100)
	}
}

func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong")
	fmt.Println("Safe operation")
}

func main() {
	err := readFile("text/output.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := []uint8{1, 2, 3, 4, 5}
	processData(data)

	safeOperation()
}
