package main

import (
	"fmt"
)

func main() {
	head := HeadCreate(2)
	isPalindrome(head.Next)
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	lastHead := slow.Next
	slow.Next = nil

	for lastHead.Next != nil {

	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func HeadCreate(n int) *ListNode {
	p := new(ListNode)
	for n > 0 {
		q := new(ListNode)
		q.Val = n
		q.Next = p.Next
		p.Next = q
		n--
	}
	return p
}

func show(head *ListNode) {
	p := head
	if p == nil {
		return
	}
	for p != nil {
		fmt.Printf("%d", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}
