package linkedlist

import (
	"testing"
)

func Test_mergeTwoSortedLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	l3 := mergeTwoSortedLists(l1, l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}
