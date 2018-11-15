package linkedlist

/*
合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l := new(ListNode)
	lp := l

	for l1 != nil || l2 != nil {
		if l1 == nil {
			lp.Next = l2
			return l.Next
		}
		if l2 == nil {
			lp.Next = l1
			return l.Next
		}

		p := new(ListNode)

		if l1.Val > l2.Val {
			p = l2
			l2 = l2.Next
		} else {
			p = l1
			l1 = l1.Next
		}

		p.Next = nil
		lp.Next = p
		lp = lp.Next
	}
	return l.Next
}
