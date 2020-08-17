package main

import "fmt"


func main() {
	postorder1 := []int{1,6,3,2,5}
	postorder2 := []int{1,3,2,6,5}
	ans1 := verifyPostorder(postorder1)
	ans2 := verifyPostorder(postorder2)
	fmt.Printf("ans1:%v, ans2:%v\n",ans1, ans2)
}

func verifyPostorder(postorder []int) bool {
    if len(postorder) <= 2 {
        return true
	}
	//后续遍历中最后一个元素是根
    root := len(postorder)-1
    for root != 0 {
        index := 0
        for postorder[index] < postorder[root] {
            index++
        }
        for postorder[index] > postorder[root] {
            index++
        }
        if root != index {
            return false
        }
        root -- 
    }
    return true
}