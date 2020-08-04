package main

import (
	"fmt"

	"github.com/aidancy/LeetCodeInGo/struct/linklist"
)

func main() {
	l := linklist.New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Print("正序：")
	l.ShowList()
	fmt.Print("\n", "反序：")
	ll := l.CallHeadNode()
	rev := reverseList(ll.HeadNode)
	for rev.Next != nil {
		fmt.Printf("\t%v", rev.Data)
		rev = rev.Next
	}
	//最后一位没有循环到
	fmt.Printf("\t%v", rev.Data)
	fmt.Print("\n")
}

func reverseList(head *linklist.Node) *linklist.Node {
	if head == nil {
		return nil
	}
	cur := head
	var pre *linklist.Node

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre

		pre = cur
		cur = tmp
	}
	return pre
}
