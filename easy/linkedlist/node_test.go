package linkedlist

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	var head ListNode
	var n1 ListNode
	n1.val = 4
	var n2 ListNode
	n2.val = 1
	var n3 ListNode
	n3.val = 7
	var n4 ListNode
	n4.val = 9
	head.Next = &n1
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	ShowAll(&head)
	//deleteNode(&n3)
	removeNthFromEnd(&head, 2)
	ShowAll(&head)
}

func ShowAll(head *ListNode) {
	p := head.Next
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}

func removeNthFromEnd(head *ListNode, n int) {
	var cur, pre *ListNode = head, head
	for n > 0 {
		pre = pre.Next
		n--
	}

	for pre.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
}
