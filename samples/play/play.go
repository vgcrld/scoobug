package main

import "fmt"

type base struct {
	a int
}

type val struct {
	base
	b string
}

func main() {
	ff := val{
		base: base{a: 1},
		b:    "2",
	}
	fmt.Println(ff.base.a, ff.b)
}
