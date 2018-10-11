package main

import (
	"fmt"
)

func main() {
	head := HeadCreate(6)
	show(head.Next)
	bl := hasCycle(head.Next)
	fmt.Println(bl)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func HeadCreate(n int) *ListNode {
	p := new(ListNode)
	i := 1
	for ; i <= n/2; i++ {
		q := new(ListNode)
		q.Val = i
		q.Next = p.Next
		p.Next = q
	}
	for ; i > 0; i-- {
		q := new(ListNode)
		q.Val = i
		q.Next = p.Next
		p.Next = q
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
