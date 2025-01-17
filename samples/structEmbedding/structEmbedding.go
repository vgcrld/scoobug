package main

import "fmt"

// base is a simple struct with a single field num.
type base struct {
	num int
}

// describe returns a string representation of the base struct.
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embeds the base struct, meaning it inherits the fields and methods of base.
// It also has an additional field str.
type container struct {
	base
	str string
}

// main is the entry point of the program. It demonstrates the use of struct embedding
// by creating an instance of container and accessing the fields and methods of the embedded base struct.
func main() {
	// Create an instance of container, initializing the embedded base struct and the str field.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Access the num field directly from the container instance.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Access the num field explicitly through the embedded base struct.
	fmt.Println("also num:", co.base.num)

	// Call the describe method from the embedded base struct.
	fmt.Println("describe:", co.describe())

	// Define an interface that includes the describe method.
	type describer interface {
		describe() string
	}

	// Assign the container instance to a variable of type describer.
	var d describer = co
	// Call the describe method through the interface.
	fmt.Println("describer:", d.describe())
}
