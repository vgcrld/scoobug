# Recover in Go

This example demonstrates the use of `recover` in Go to handle panics. The `recover` function regains control of a panicking goroutine, while `defer` is used to call the `recover` function.

## Code

```go
package main

import (
	"fmt"
)

// main is the entry point of the program. It demonstrates the use of recover to handle a panic
// caused by a division by zero in the divideTwoNumbers function.
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

// divideTwoNumbers divides two integers and returns a struct containing the result and any error message.
// It uses defer and recover to handle a potential division by zero panic.
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
```

## Explanation

### main Function

The `main` function is the entry point of the program. It demonstrates the use of `recover` to handle a panic caused by a division by zero in the `divideTwoNumbers` function.

- The `main` function starts by printing a message.
- It calls the `divideTwoNumbers` function with arguments `10` and `0`, which causes a division by zero panic.
- The result of the division is printed.
- The `main` function ends by printing a message.
- A record is created and printed.

### divideTwoNumbers Function

The `divideTwoNumbers` function divides two integers and returns a struct containing the result and any error message. It uses `defer` and `recover` to handle a potential division by zero panic.

- The `defer` statement is used to call an anonymous function that calls `recover`.
- If a panic occurs, `recover` regains control and returns the panic value.
- The error message is stored in the `err` field of the result struct.
- The division is performed, and the result is stored in the `result` field of the result struct.
- The result struct is returned.
