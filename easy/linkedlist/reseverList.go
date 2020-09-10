package linkedlist

/*
反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 迭代
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head.Next

	// 类似头插法
	for p.Next != nil {
		q := p.Next
		p.Next = q.Next
		q.Next = head.Next
		head.Next = q
	}

	p.Next = head // 形成环
	head = head.Next
	p.Next.Next = nil // 断掉环

	return head
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head
	var prev *ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

/** 递归 C:
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 * struct ListNode* reverseList(struct ListNode* head) {
 *    if (head == NULL || head->next == NULL) return head;
 *
 *    struct ListNode* newHead = reverseList(head->next);
 *    head->next->next = head;
 *    head->next = NULL;
 *    return newHead;
 *}
 */
