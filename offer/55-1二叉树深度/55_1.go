package depth

import "github.com/aidancy/LeetCodeInGo/struct/binarytree"

//MaxDepth dfs递归
func MaxDepth(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Lchild)
	right := MaxDepth(root.Rchild)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//MaxDepth2 dfs回溯法
func MaxDepth2(root *binarytree.TreeNode) int {
	maxInt := 0
	count := 0
	var helper func(root *binarytree.TreeNode)
	helper = func(root *binarytree.TreeNode) {
		if root == nil {
			if count > maxInt {
				maxInt = count
			}
			return
		}
		count++
		helper(root.Lchild)
		helper(root.Rchild)
		count--
	}
	helper(root)
	return maxInt
}

//MaxDepth3 bfs回溯法
func MaxDepth3(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*binarytree.TreeNode{root}
	maxInt := 0
	for len(q) > 0 {
		maxInt++
		l := len(q)
		for l > 0 {
			if q[0].Lchild != nil {
				q = append(q, q[0].Lchild)
			}
			if q[0].Rchild != nil {
				q = append(q, q[0].Rchild)
			}
			q = q[1:]
			l--
		}
	}
	return maxInt
}
