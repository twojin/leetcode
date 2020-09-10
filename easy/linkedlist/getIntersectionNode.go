package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currA, currB := headA, headB

	for currA != currB {
		if currA == nil {
			currA = headB
		} else {
			currA = currA.Next
		}

		if currB == nil {
			currB = headA
		} else {
			currB = currB.Next
		}
	}

	return currA
}
