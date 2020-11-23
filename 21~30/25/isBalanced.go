func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if absDiff(lh, rh) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func absDiff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxHeight(height(root.Left), height(root.Right))
}

func maxHeight(a, b int) int {
	if a >= b {
		return a
	}
	return b
}