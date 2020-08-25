package tree

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}

func getMedian(left, right *ListNode) *ListNode {
	slow := left
	fast := left

	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
