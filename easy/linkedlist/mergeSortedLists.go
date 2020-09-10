package linkedlist

func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	ln := l3
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			ln.Next = l2
			l2 = l2.Next
		} else {
			ln.Next = l1
			l1 = l1.Next
		}
		ln = ln.Next
	}

	if l1 != nil {
		ln.Next = l1
	}

	if l2 != nil {
		ln.Next = l2
	}

	return l3.Next
}
