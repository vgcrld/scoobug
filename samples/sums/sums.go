package main

import "fmt"

type Result struct {
	Index1 int
	Index2 int
	Value  int
}

func inspectSlice(nums []int, target int) (ret []Result) {
	for i1, v1 := range nums {
		for i2, v2 := range nums {
			ret = append(ret, Result{i1, i2, v1 + v2})
		}
	}
	return returnAnswer(ret, target)
}

func returnAnswer(r []Result, p int) (ret []Result) {

	for _, res := range r {
		if res.Value == p {
			ret = append(ret, res)
		}
	}

	return ret

}

func main() {

	xx := inspectSlice([]int{11, 2, 8, 4, 5, 6, 6}, 13)

	fmt.Println(xx)

}
