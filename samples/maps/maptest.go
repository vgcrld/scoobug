package main

import (
	"fmt"
	"os"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	x := strings.Fields(s)
	for _, v := range x {
		v = strings.ToLower(v)
		if _, ok := ret[v]; ok {
			ret[v] = ret[v] + 1
		} else {
			ret[v] = 1
		}
	}
	return ret
}

func PrintMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("%-15s %4v\n", k, v)
	}
}

func main() {
	s := strings.Join(os.Args[1:], " ")
	wc := WordCount(s)
	PrintMap(wc)
}
