package main

import (
	"fmt"
)

func main() {

	// Defer a function to handle panic recovery
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Example function that will cause a panic
	funcThatPanics()

	fmt.Println("This line will not be executed due to panic")
}

// funcThatPanics demonstrates a function that causes a panic
func funcThatPanics() {
	defer fmt.Println("Deferred call in funcThatPanics")
	fmt.Println("About to panic")
	panic("Something went wrong!")
}
