package main

import (
	"fmt"

	"github.com/aidancy/LeetCodeInGo/struct/binarytree"
)

func main() {
	root := binarytree.New(3)
	a := binarytree.New(1)
	b := binarytree.New(2)
	d := binarytree.New(4)
	e := binarytree.New(5)
	root.Lchild = d
	root.Rchild = e
	d.Lchild = a
	d.Rchild = b
	sub := binarytree.New(4)
	f := binarytree.New(1)
	sub.Lchild = f
	ans := isSubStructure(root, sub)
	fmt.Printf("是否是字结构？：%v\n", ans)
}

func isSubStructure(A *binarytree.TreeNode, B *binarytree.TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var ret bool

	if A.Data == B.Data {
		ret = check(A, B)
	}

	if !ret {
		ret = isSubStructure(A.Lchild, B)
	}
	if !ret {
		ret = isSubStructure(A.Rchild, B)
	}
	return ret
}

func check(A *binarytree.TreeNode, B *binarytree.TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Data != B.Data {
		return false
	}
	//递归校验 A B 左子树和右子树的结构和节点是否相同
	return check(A.Lchild, B.Lchild) && check(A.Rchild, B.Rchild)
}
