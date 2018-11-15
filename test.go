package main

import (
	"fmt"
)

func main() {
}

func climbStairs(n int) int {
	if n <= 0 || n == 1 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	t2t := make([]int, 0)
	midTraversal(root, t2t)

	for i := 1; i < len(t2t); i++ {
		if t2t[i] <= t2t[i-1] {
			return false
		}
	}

	return true
}

func midTraversal(root *TreeNode, t2t []int) {
	if root == nil {
		return
	}

	midTraversal(root.Left, t2t)
	t2t = append(t2t, root.Val)
	midTraversal(root.Right, t2t)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func HeadCreate(n int) *ListNode {
	p := new(ListNode)
	i := 1
	for ; i <= n/2; i++ {
		q := new(ListNode)
		q.Val = i
		q.Next = p.Next
		p.Next = q
	}
	for ; i > 0; i-- {
		q := new(ListNode)
		q.Val = i
		q.Next = p.Next
		p.Next = q
	}
	return p
}

func show(head *ListNode) {
	p := head
	if p == nil {
		return
	}
	for p != nil {
		fmt.Printf("%d", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}
