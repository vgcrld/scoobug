package main

// Recover is a built-in function that regains control of a panicking goroutine.
// Defer is used to call recover function to regain control of a panicking goroutine.
// Panic is a built-in function that stops the ordinary flow of control and begins panicking.

import (
	"fmt"
)

func main() {
	fmt.Println("Start of main function")
	val := divideTwoNumbers(10, 0)
	fmt.Println(val.result)
	fmt.Println("End of main function")
	record := struct {
		Name string
		Age  int
	}{"John", 30}
	fmt.Println(record)
}

func divideTwoNumbers(a, b int) (res struct {
	err    string
	result int
}) {

	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println("Recovered in divideTwoNumbers:", rec)
			res.err = rec.(error).Error()
			// res.result = -1
		}
	}()

	res.result = a / b
	return res
}
