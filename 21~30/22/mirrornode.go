func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, root.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
	} else {
		return p == q
	}
}
