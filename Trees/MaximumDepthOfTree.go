package Trees

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := MaxDepth(root.Left)
	maxRight := MaxDepth(root.Right)
	return max(maxLeft, maxRight) + 1
}
