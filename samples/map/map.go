package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["key1"] = 10
	x["key2"] = 11
	x["key3"] = 12

	fmt.Println(x["key1"])
	fmt.Println(x["key3"])
	fmt.Println(x["key2"])
}
