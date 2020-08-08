package tree

import (
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	var tree TreeNode
	tree.Val = 1
	tree.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	tree.Right = &TreeNode{Val: 3}

	t.Log(diameterOfBinaryTree(&tree))
}
