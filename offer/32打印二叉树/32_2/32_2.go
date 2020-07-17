package main

import (
	"fmt"

	"github.com/Aidanwangcy/leetcode/struct/binarytree"
)

func main()  {
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
	//初始条件，插入根，为第0层
	queue := []*binarytree.TreeNode{root}
	level := 0

	for len(queue) != 0 {
		//临时队列，为了插入下一层的节点
		temp := []*binarytree.TreeNode{}
		//每执行一次循环添加一层数组
		res = append(res, []interface{}{})
		for _, v := range queue {
			//给当前层的数组添加元素
			res[level] = append(res[level], v.Data)
			//顺便把下一层的节点加到临时数组里面
			if v.Lchild != nil {
				temp = append(temp, v.Lchild)
			}
			if v.Rchild != nil {
				temp = append(temp, v.Rchild)
			}
		}
		//循环完当前层，层数+1
		level++
		//把下一层的节点更换到队列里
		queue =temp
	}
	return res
}