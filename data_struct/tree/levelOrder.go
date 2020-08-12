package tree

func levelOrder(root *TreeNode) [][]int {
	output := [][]int{}
	if root == nil {
		return output
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		count := len(queue)
		level := make([]int, 0, count)

		for count > 0 {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			count--
		}
		output = append(output, level)
	}

	return output
}
