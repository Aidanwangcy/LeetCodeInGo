package binarytree

import (
	"container/list"
	"fmt"
)

//TreeNode 二叉树节点
type TreeNode struct {
	Data interface{}
	Lchild *TreeNode
	Rchild *TreeNode
}

//Initer 创建新节点
type Initer interface {
    SetData (data interface{})
}

//Operater 二叉树输出，深度计算，叶子节点统计
type Operater interface {
    PrintBT()
    Depth() int
    LeafCount() int
}

//Order 遍历
type Order interface {
    PreOrder()
    InOrder()
    PostOrder()
}

//New 构造
func New(data interface{}) *TreeNode {
	return &TreeNode{Data: data}
}

//SetData 节点赋值
func (node *TreeNode) SetData(data interface{}) {
	node.Data = data
}

//PrintBT 输出二叉树括号表示
func (node *TreeNode) PrintBT() {
	PrintBT(node)
}

//LeafCount 返回叶子节点数
func (node *TreeNode) LeafCount() int {
	return LeafCount(node)
}

//PreOrder 前序遍历
func (node *TreeNode) PreOrder() {
	PreOrder(node)
}

//InOrder 前序遍历
func (node *TreeNode) InOrder() {
	InOrder(node)
}

//PostOrder 前序遍历
func (node *TreeNode) PostOrder() {
	PostOrder(node)
}

//层序遍历

//DFS 深度优先
func (node *TreeNode) DFS() [][]interface{} {
	return DFS(node)
}

//BFS 广度优先
func (node *TreeNode) BFS() [][]interface{} {
	return BFS(node)
}


func PrintBT(node *TreeNode) {
	if node != nil {
		fmt.Printf("%v", node.Data)
		if node.Lchild != nil || node.Rchild != nil {
			fmt.Printf("(")
			PrintBT(node.Lchild)
			if node.Rchild != nil {
				fmt.Printf(",")
			}
			PrintBT(node.Rchild)
			fmt.Printf(")")
		}
	}
}

func Depth(node *TreeNode) int {
	var depl, depr int
	if node == nil {
		return 0
	}
	depl = Depth(node.Lchild)
	depr = Depth(node.Rchild)
	if depl > depr {
		return depl + 1
	}
	return depr + 1
}

func LeafCount(node *TreeNode) int {
	if node == nil {
		return 0
	} else if (node.Lchild == nil) && (node.Rchild == nil) {
		return 1
	} else {
		return (LeafCount(node.Lchild) + LeafCount(node.Rchild))
	}
}

func PreOrder(node *TreeNode) {
	if node != nil {
		fmt.Printf("%v", node.Data)
		PreOrder(node.Lchild)
		PreOrder(node.Rchild)
	}
}

func InOrder(node *TreeNode) {
	if node != nil {
		InOrder(node.Lchild)
		fmt.Printf("%v", node.Data)
		InOrder(node.Rchild)
	}
}

func PostOrder(node *TreeNode) {
	if node != nil {
		PreOrder(node.Lchild)
		PreOrder(node.Rchild)
		fmt.Printf("%v", node.Data)
	}
}

func DFS(node *TreeNode) [][]interface{} {
	var res [][]interface{}
	var _dfs func(node *TreeNode, level int)
	_dfs = func(node *TreeNode, level int){
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, []interface{}{})
		}
		res[level] = append(res[level], node.Data)

		_dfs(node.Lchild, level + 1)
		_dfs(node.Rchild, level + 1)
	}

	_dfs(node, 0)
	return res
}

func BFS(node *TreeNode) [][]interface{} {
	var res [][]interface{}
	if node == nil {
		return res
	}

	list := list.New()
	list.PushFront(node)
	for list.Len() > 0 {
		var currentLevel []interface{}
		listLength := list.Len()
		for i := 0; i < listLength; i++ {
			LastNode := list.Remove(list.Back()).(*TreeNode)
			currentLevel = append(currentLevel, LastNode.Data)
			if LastNode.Lchild != nil {
				list.PushFront(LastNode.Lchild)
			}
			if LastNode.Rchild != nil {
				list.PushFront(LastNode.Rchild)
			}
		}
		res = append(res, currentLevel)
	}
	return res
}