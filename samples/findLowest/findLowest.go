package main

import (
	"fmt"
	"os"
	"strconv"
)

// Bubble sort - return the lowest number in a slice
func findLowest(s []int) int {
	r := s[0]
	for _, i := range s {
		// fmt.Println(i)
		if i < r {
			r = i
		}
	}

	return r
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: findLowest <number> <number> ...")
		os.Exit(1)
	}

	intArgs := parseArgs(os.Args[1:])
	l := findLowest(intArgs)
	fmt.Println("Lowest number is:", l)
}

func parseArgs(a []string) []int {
	var r []int
	for _, i := range a {
		i, _ := strconv.Atoi(i)
		r = append(r, i)
	}
	return r
}
