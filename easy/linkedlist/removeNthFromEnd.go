package linkedlist

/*
删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre, cur := head, head
	for n > 0 { // 使用双指针 一个指针先行n步
		pre = pre.Next
		n--
	}

	if pre == nil { // 删除的是链表的头部
		return head.Next
	}

	for pre.Next != nil { // 当先行指针下一个节点为空时 当前指指针的下一个节点就是倒数第n个节点
		cur = cur.Next
		pre = pre.Next
	}

	cur.Next = cur.Next.Next
	return head
}

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 *
 * struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
 *	struct ListNode* cur = head;
 *	struct ListNode* pre = head;
 *
 *	while(n){
 *		pre = pre->next;
 *		n--;
 *	}
 *
 *	if (pre == NULL) return head->next;
 *
 *	while(pre->next){
 *		pre = pre->next;
 *		cur = cur->next;
 *	}
 *	cur->next = cur->next->next;
 *	return head;
 *}
 */
