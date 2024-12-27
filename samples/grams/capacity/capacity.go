package main

import (
	"fmt"
	"math"
)

type Capacity float64

func (u Capacity) Convert() string {
	var cap Capacity
	var mod string
	switch {
	case u >= 1<<50:
		cap = u / (1 << 50)
		mod = "Pi"
	case u >= 1<<40:
		cap = u / (1 << 40)
		mod = "Ti"
	case u >= 1<<30:
		cap = u / (1 << 30)
		mod = "Gi"
	case u >= 1<<20:
		cap = u / (1 << 20)
		mod = "Mi"
	case u >= 1<<10:
		cap = u / (1 << 10)
		mod = "Ki"
	default:
		cap = u
		mod = "Bi"
	}
	return fmt.Sprintf("%.1f %s", cap, mod)
}

func main() {

	x := math.Pow(1024, 5) * 1234
	fmt.Println(Capacity.Convert(Capacity(x)))

}
