package linkedlist

/*
环形链表
给定一个链表，判断链表中是否有环。

进阶：
你能否不使用额外空间解决此题？

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 ******************** C:
 * bool hasCycle(struct ListNode *head) {
 *    if (head == NULL) return false;
 *    struct ListNode *fast = head;
 *    struct ListNode *slow = head;
 *
 *    while(fast->next != NULL && fast->next->next != NULL){
 *        slow = slow->next;
 *        fast = fast->next->next;
 *        if (fast == slow) return true;
 *    }
 *
 *    return false;
 *}
**/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
