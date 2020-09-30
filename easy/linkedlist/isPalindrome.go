package linkedlist

/*
回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	p := slow.Next
	for p.Next != nil {
		q := p.Next
		p.Next = q.Next
		q.Next = slow.Next
		slow.Next = q
	}
	fastHead := slow.Next
	slow.Next = nil

	for fastHead != nil {
		if fastHead.Val != head.Val {
			return false
		}
		fastHead = fastHead.Next
		head = head.Next
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到mid节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 翻转后半部分链表
	slow = slow.Next
	var dummy *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = dummy
		dummy = slow
		slow = next
	}

	for dummy != nil {
		if dummy.Val != head.Val {
			return false
		}
		dummy, head = dummy.Next, head.Next
	}

	return true
}
