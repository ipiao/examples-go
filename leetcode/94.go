package leetcode

func inorderTraversal(root *TreeNode) []int {
	var leftstack []*TreeNode

	lnode := root
	for lnode != nil {
		leftstack = append(leftstack, lnode)
		lnode = lnode.Left
	}
	var res []int
	for l := len(leftstack); l > 0; {
		cur := leftstack[l-1]
		res = append(res, cur.Val)
		leftstack = leftstack[:l-1]
		if cur.Right != nil {
			lnode = cur.Right
			for lnode != nil {
				leftstack = append(leftstack, lnode)
				lnode = lnode.Left
			}
		}
	}
	return res
}
