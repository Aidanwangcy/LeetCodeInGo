package main

import (
	"fmt"
	"math"
)

//TreeNode 二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	ans1 := maxDepth(root)
	ans2 := maxDepth2(root)

	fmt.Printf("二叉树最大深度为：%v\n", ans1)
	fmt.Printf("2.二叉树最大深度为：%v\n", ans2)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return int(math.Max(float64(left), float64(right)) + 1)
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode       // 辅助队列
	queue = append(queue, root) // 根节点入队
	depth := 0                  // 初始化深度为0

	for len(queue) > 0 { // 当队列不为空时，循环以下操作
		size := len(queue)          //当前队列中元素个数size作为限定:当前层级中节点数量
		for i := 0; i < size; i++ { // 遍历当前层级中的节点
			s := queue[0]      // 获取队首元素
			queue = queue[1:]  // 队首元素出队
			if s.Left != nil { // 如果左子树不为空，左子树入队
				queue = append(queue, s.Left)
			}
			if s.Right != nil { // 如果右子树不为空，右子树入队
				queue = append(queue, s.Right)
			}
		}
		depth++ // for循环结束后这一层级节点已经访问结束，深度加+1
	}
	return depth
}
