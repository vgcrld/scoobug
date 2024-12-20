package main

import "fmt"

func main() {

	var x int = 10
	fmt.Printf("Original: %v\n", x)
	changeValue(&x)
	fmt.Printf("After Change: %v\n", x)

}

func changeValue(x *int) {
	*x = 20
}
