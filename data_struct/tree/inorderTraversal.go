package tree

/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xecaj6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func inorderTraversal(root *TreeNode) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}

	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = Pop(&stack)
		output = append(output, curr.Val)
		curr = curr.Right
	}

	return output
}
