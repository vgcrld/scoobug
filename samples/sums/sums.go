package main

import (
	"fmt"
	"os"
	"strconv"
)

// Printer interface defines a method for printing
type Printer interface {
	String()
}

// Result struct holds three integers
type Result struct {
	x int
	y int
	z int
}

// String method prints the Result struct
func (r Result) String() {
	fmt.Printf("%d,%d\n", r.x, r.y)
}

// main function
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: sums <number> [<num1> <num2> ...]")
		os.Exit(1)
	}

	pick, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error converting pick number to integer:", err)
		os.Exit(1)
	}

	os.Args = os.Args[1:]

	nums := buildIt(getArgs(), pick)
	for _, r := range nums {
		r.String()
	}

}

// getArgs returns a slice of integers from command line arguments
func getArgs() []int {
	args := os.Args[1:]
	var intArgs []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error converting argument to integer:", err)
			os.Exit(1)
		}
		intArgs = append(intArgs, num)
	}
	return intArgs
}

// buildIt returns a slice of Result structs
func buildIt(nums []int, pick int) (ret []Result) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == pick {
				ret = append(ret, Result{i, j, sum})
			}
		}
	}
	return ret
}
