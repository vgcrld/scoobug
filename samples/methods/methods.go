package main

import "fmt"

// rect represents a rectangle with a width and height.
type rect struct {
	width, height int
}

// area calculates the area of a rectangle. It uses a pointer receiver to modify the original rect.
func (r *rect) area() int {
	return r.width * r.height
}

// perim calculates the perimeter of a rectangle. It uses a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// main is the entry point of the program. It creates an instance of rect and demonstrates
// calling methods with both value and pointer receivers.
func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
