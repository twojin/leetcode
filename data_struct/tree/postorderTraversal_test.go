package tree

import (
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	/*
		     1
		  2     3
		4  5

	*/

	var tree TreeNode
	tree.Val = 1
	tree.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	tree.Right = &TreeNode{Val: 3}

	t.Log(postorderTraversal(&tree))
}
