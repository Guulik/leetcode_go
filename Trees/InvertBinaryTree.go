package Trees

func InvertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	inverted := root
	if root.Left != nil {
		InvertBinaryTree(root.Left)
	}
	if root.Right != nil {
		InvertBinaryTree(root.Right)
	}
	inverted.Left, inverted.Right = root.Right, root.Left
	return inverted
}
