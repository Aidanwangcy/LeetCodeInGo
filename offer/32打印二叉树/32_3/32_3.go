package main

import (
	"fmt"

	"github.com/Aidanwangcy/leetcode/struct/binarytree"
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
		for _, v1 := range v {
			fmt.Printf("%v ", v1)
		}
		fmt.Printf("\n")
	}
}

func levelOrder(root *binarytree.TreeNode) [][]interface{} {
	var res [][]interface{}
	if root == nil {
		return res
	}

	queue := []*binarytree.TreeNode{root}
	level := 0
	for len(queue) != 0 {
		temp := []*binarytree.TreeNode{}
		res = append(res, []interface{}{})
		for _, v := range queue {
			if level % 2 == 0 {
				res[level] = append(res[level], v.Data)
			} else {
				restemp := []interface{}{v.Data}
				res[level] = append(restemp,res[level]...)
			}
			if v.Lchild != nil {
				temp = append(temp,v.Lchild)
			}
			if v.Rchild != nil {
				temp = append(temp, v.Rchild)
			}
		}
		queue = temp
		level ++
	}
	return res
}
