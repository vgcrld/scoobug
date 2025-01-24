package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String("name", "Rich Davis", "User's name")
	age := flag.Int("age", 0, "User's age")
	print := flag.Bool("print", false, "Print the name and age")

	flag.Parse()

	if *print {
		fmt.Printf("Name: %s, Age: %d\n", *name, *age)
	}
	fmt.Println("Flags:", flag.Args())

}
