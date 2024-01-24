/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	longestPath(root, &diameter)
	return diameter
}

func longestPath(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}

	leftPath := longestPath(node.Left, diameter)
	rightPath := longestPath(node.Right, diameter)

	*diameter = max(*diameter, leftPath+rightPath)

	return max(leftPath, rightPath) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}