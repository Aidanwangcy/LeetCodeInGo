package main

import (
	"fmt"

	"github.com/Aidanwangcy/LeetCodeInGo/struct/binarytree"
)

func main() {
	root := binarytree.New(3)
	a := binarytree.New(9)
	b := binarytree.New(20)
	c := binarytree.New(15)
	d := binarytree.New(7)
	b.Lchild = c
	b.Rchild = d
	root.Lchild = a
	root.Rchild = b
	ans := levelOrder(root)
	for _, v := range ans {
		fmt.Printf("%v ", v)
	}
}

func levelOrder(root *binarytree.TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	
	var ret []interface{}

	//根入队
	queue := []*binarytree.TreeNode{root}

	//队列不为空
	for len(queue) != 0 {
		temp := []*binarytree.TreeNode{}
		for _, v := range queue {
			ret = append(ret, v.Data)
			if v.Lchild != nil {
				temp = append(temp, v.Lchild)
			}
			if v.Rchild != nil {
				temp = append(temp, v.Rchild)
			}
		}
		//下一层
		queue = temp
	}
	return ret
}
