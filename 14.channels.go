package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channels are a way to communicate between go routines
// they:
// 1. hold data - like a variable
// 2. thread safe - like a mutex, they help us avoid data race when multiple go routines are trying to access the same data
// 3. listen for data - like a listener, they can wait for data to be sent to them and then do something with it

var MAX_PHONE_PRICE float32 = 1000.0

func main() {
	var phoneChannel = make(chan string)                                                      // create a channel that holds strings
	var carriers = []string{"Pelephone", "Cellcom", "Partner", "Hot Mobile", "Golan Telecom"} // list of carriers
	for i := range carriers {
		go checkPhonePrice(carriers[i], phoneChannel) // start a go routine for each carrier to check the phone price in the background
	}
	sendMessage(phoneChannel) // wait for a message from the phone channel
}

func checkPhonePrice(carrier string, phoneChannel chan string) {
	for {
		time.Sleep(time.Second * 1)        // wait for a second
		phonePrice := rand.Float32() * 20  // generate a random price for the phone
		if phonePrice <= MAX_PHONE_PRICE { // if the price is less than the max price
			phoneChannel <- carrier // send a message to the phone channel
			break                   // stop the loop
		}
	}
}

func sendMessage(phoneChannel chan string) {
	fmt.Printf("\nFound a deal on phone at %s", <-phoneChannel) // wait for a message from the phone channel and print it
}
