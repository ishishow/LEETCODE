func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return diveNode(root.Left, root.Right, 1)
}

func diveNode(p *TreeNode, q *TreeNode, count int) int {
	if p != nil && q != nil {
		return max(diveNode(p.Left, p.Right, count+1), diveNode(q.Left, q.Right, count+1))
	} else if p != nil {
		return diveNode(p.Left, p.Right, count+1)
	} else if q != nil {
		return diveNode(q.Left, q.Right, count+1)
	} else {
		return count
	}
}

func max(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret < v {
			ret = v
		}
	}
	return ret
}