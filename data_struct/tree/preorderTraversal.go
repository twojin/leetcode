package tree

/*
二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？


Go



作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xeywh5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tree []int

func preorderTraversal(root *TreeNode) []int {
	tree = make([]int, 0, 0)
	if root == nil {
		return tree
	}

	preorder(root)
	return tree
}

func preorder(node *TreeNode) {
	if node == nil {
		return
	}

	tree = append(tree, node.Val)
	preorder(node.Left)
	preorder(node.Right)
}

func traversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 0)
	output := make([]int, 0, 0)
	if root == nil {
		return output
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		node := pop(&stack)
		output = append(output, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return output
}

func pop(stack *[]*TreeNode) *TreeNode {
	if len(*stack) < 1 {
		return nil
	}

	node := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return node
}
