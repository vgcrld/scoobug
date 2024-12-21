package main

import (
	"fmt"
	"math"
)

// geometry is an interface that defines methods for calculating area and perimeter.
type geometry interface {
	area() float64
	perim() float64
}

// rect represents a rectangle with a width and height.
type rect struct {
	width, height float64
}

// circle represents a circle with a radius.
type circle struct {
	radius float64
}

// area calculates the area of a rectangle.
// the formula for the area of a rectangle is width * height
func (r rect) area() float64 {
	return r.width * r.height
}

// perim calculates the perimeter of a rectangle.
// the formula for the perimeter of a rectangle is 2 * (width + height)
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// area calculates the area of a circle.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// perim calculates the perimeter of a circle.
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// measure takes a geometry interface and prints its details, area, and perimeter.
func measure(g geometry) {
	a := g.area()
	p := g.perim()
	switch g := g.(type) {
	case rect:
		fmt.Printf("Rectangle width(%v), hight(%v): area=%.2f, perim=%.2f\n", g.width, g.height, a, p)
	case circle:
		fmt.Printf("Circle radius(%v): area=%.2f, perim=%.2f\n", g.radius, a, p)
	default:
		fmt.Println("Unknown")
	}
}

// main is the entry point of the program. It creates instances of rect and circle,
// and then calls the measure function with these instances to demonstrate polymorphism
// with interfaces.
func main() {
	r := rect{width: 3.1, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
