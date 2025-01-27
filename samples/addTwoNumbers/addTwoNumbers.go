package main

import (
	"log"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// ListNode defines a node in the singly linked list
type ListNode struct {
	Val  int
	Name string
	Next *ListNode
}

func (l *ListNode) list() (ret []ListNode) {
	for l != nil {
		ret = append(ret, *l)
		l = l.Next
	}
	return
}

func (l *ListNode) maps() map[string]*ListNode {
	ret := make(map[string]*ListNode)
	for l != nil {
		ret[l.Name] = l
		l = l.Next
	}
	return ret
}

func main() {

	// Create a slice of integers
	nums := []int{100, 1, 2, 2000}

	// Build the linked list from the slice
	ll := &ListNode{Val: nums[0]}
	num := nums[0]
	nnum := strconv.Itoa(num)
	log.Println("Working on 'num': ", num)
	current := ll
	current.Val = num
	current.Name = "name" + nnum
	for _, num := range nums[1:] {
		log.Println("Working on 'num': ", num)
		current.Next = &ListNode{Val: num}
		current = current.Next
		nnum := strconv.Itoa(num)
		current.Name = "name" + nnum
		current.Val = num
	}

	spew.Dump(ll)

}
