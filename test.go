package main

import (
	"fmt"
)

func main() {

	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}
	fmt.Println(isPalindrome(node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
