package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode 
}

func main() {
	var head ListNode
	var node1 ListNode
	node1.Val = 1
	var node2 ListNode
	node2.Val = 3
	var node3 ListNode
	node2.Val = 2
	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3

	point := head.Next
	for point != nil {
		fmt.Print(point.Val)
		point = point.Next
	}
	ans := reversePrint(&head)
	fmt.Println(ans)
}

func reversePrint(head *ListNode) []int {
	var queue1, queue2 []int
	point := head.Next
	
	for point != nil {
		queue1 = append(queue1, point.Val)
		point = point.Next
	}

	for j := len(queue1)-1; j >= 0 ; j-- {
		queue2 = append(queue2, queue1[j])
	}
	return queue2
}