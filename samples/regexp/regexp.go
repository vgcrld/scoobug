package main

import (
	"fmt"
	"regexp"
)

func main() {

	// Compile a regular expression pattern to match a string in the format "name[number]".
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))

	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	if validID.MatchString(input) {
		fmt.Println("Matched!")
	} else {
		fmt.Println("Not matched!")
	}

}
