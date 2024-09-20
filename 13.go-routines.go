package main

// Concurrency is the ability to run multiple programs at the same time
// Concurrency != Parallelism
// Concurrency is about dealing with lots of things at once
// Parallelism is about doing lots of things at once

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}    // Read Write Mutex
var wg = sync.WaitGroup{} // Wait Group to wait for all the go routines to finish
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	// mainTask()     // Concurrent execution avg time: 2s
	// constantTime() // Concurrent execution avg time: 2s
	linearTime() // Concurrent execution avg time: 2s
}

func dbCall(i int) {
	// simulate a database call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	saveToResults(dbData[i]) // Saving the result to the results slice
	log()                    // Logging the results
	wg.Done()                // Removing the go routine from the wait group
}

func dbCall2(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	wg.Done()
}

func count() {
	// simulate a long running process
	var res int
	for i := 0; i < 1000000000; i++ {
		res += i
	}
	wg.Done()
}

func saveToResults(data string) {
	m.Lock()                        // Locking the memory location
	results = append(results, data) // this alone will cause multiple go routines to write to the same memory location at the same time
	m.Unlock()                      // Unlocking the memory location after writing to it
}

func log() {
	m.RLock()                                    // Locking the memory location for reading
	fmt.Printf("\nThe results are: %v", results) // Reading from the memory location
	m.RUnlock()                                  // Unlocking the memory location after reading from it
}

func mainTask() {
	t0 := time.Now() // Start time
	for i := 0; i < len(dbData); i++ {
		// dbCall(i) // Sequential execution avg time: 10s
		wg.Add(1)    // Adding the go routine to the wait group
		go dbCall(i) // Concurrent execution avg time: 2s
	}
	wg.Wait() // Waiting for all the go routines to finish
	fmt.Printf("\nThe results are: %v", results)
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0)) // Total execution time: 10.0000001s
}

func constantTime() {
	t0 := time.Now()
	// this will take 2s to execute because the go routines are not doing anything
	// they are just waiting for 2s and then they are done
	for i := 0; i < 1000; i++ { // 2s execution time
		wg.Add(1)
		go dbCall2(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
}

func linearTime() {
	t0 := time.Now()
	// this will take more than 2s to execute because the go routines are doing something
	// they are counting from 0 to 1000000000
	// so the execution time will be more than 2s
	// 1 go routine takes 530ms to execute
	// 10 go routines take 846ms
	// 100 go routines take 8.6s
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\nlinearTime total execution time: %v\n", time.Since(t0))
}
