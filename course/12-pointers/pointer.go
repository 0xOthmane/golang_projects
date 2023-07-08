package main

import (
	"fmt"
	"strings"
)

/* 
The * syntax defines a pointer:
	var p *int
A pointer's zero value is nil
The & operator generates a pointer to its operand.
	myString := "hello"
	myStringPtr = &myString
The * dereferences a pointer to gain access to the value

	fmt.Println(*myStringPtr) // read myString through the pointer
	*myStringPtr = "world"    // set myString through the pointer 
*/

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}