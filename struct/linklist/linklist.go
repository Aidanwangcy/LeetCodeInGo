package linklist

//用在offer24题
import (
	"errors"
	"fmt"
)

//Object 数据域
type Object interface{}

//Node 节点组成
type Node struct {
	Data Object //数据域
	Next *Node  //地址域：下一个节点
}

//HeadLinkNode 头节点
type HeadLinkNode struct {
	HeadNode *Node
}

//LinkNoder 定义节点接口
type LinkNoder interface {
	IsEmpty() bool
	Len() int
	Add(data Object) *Node
	Append(data Object)
	Insert(data Object, index int)
	Remove(data Object)
	RemoveAtInsex(index int) error
	Contain(data Object) bool
	ShowList()
	CallHeadNode() *HeadLinkNode
}

//New 新建头节点
func New() LinkNoder {
	return &HeadLinkNode{HeadNode: &Node{Data: "head", Next: nil}}
}

//CallHeadNode 返回头节点
func (list *HeadLinkNode) CallHeadNode() *HeadLinkNode {
	return list
}

//IsEmpty 判断是否为空
func (list *HeadLinkNode) IsEmpty() bool {
	if list.HeadNode == nil {
		return true
	}
	return false
}

//Len 获取列表长度
func (list *HeadLinkNode) Len() int {
	cur := list.HeadNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

//Add 链表头部添加元素
func (list *HeadLinkNode) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = list.HeadNode
	list.HeadNode = node
	return node
}

//Append 尾部添加元素
func (list *HeadLinkNode) Append(data Object) {
	node := &Node{Data: data}
	if list.IsEmpty() {
		list.HeadNode = node
	} else {
		cur := list.HeadNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

//Insert 插入
func (list *HeadLinkNode) Insert(data Object, index int) {
	if index < 0 {
		list.Add(data)
	} else if index > list.Len() {
		list.Append(data)
	} else {
		pre := list.HeadNode
		count := 0
		for count < index-1 {
			pre = pre.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

//Remove 删除某个值的节点
func (list *HeadLinkNode) Remove(data Object) {
	pre := list.HeadNode
	if pre.Data == data {
		list.HeadNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

//RemoveAtInsex 删除制定位置的节点
func (list *HeadLinkNode) RemoveAtInsex(index int) error {
	pre := list.HeadNode
	if index <= 0 { //删除头节点
		list.HeadNode = pre.Next
	} else if index > list.Len() {
		return errors.New("超出链表长度")
	} else {
		count := 0
		for count != (index-1) && pre.Next != nil {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
	return nil
}

//Contain 查看是否包含某个元素
func (list *HeadLinkNode) Contain(data Object) bool {
	cur := list.HeadNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

//ShowList 遍历
func (list *HeadLinkNode) ShowList() {
	if !list.IsEmpty() {
		cur := list.HeadNode
		for {
			fmt.Printf("\t%v", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	} else {
		fmt.Print("空")
	}
}
