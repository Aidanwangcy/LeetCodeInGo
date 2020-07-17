package main

import "fmt"

func main()  {
	pushed := []int{1,2,3,4,5}
	popped := []int{4,5,3,2,1}
	ans := validateStackSequences(pushed, popped)
	fmt.Printf("%v\n", ans)
}

func validateStackSequences(pushed []int, popped []int) bool {
	//如果入栈序列长度为0，一定为真
	if len(pushed) == 0 {
		return true
	}
	//辅助slice，实现栈的操作
	temp := []int{}
	//计数每取出一个+1
	i := 0
	//pushed放入临时栈
	for _, v := range pushed {
		temp = append(temp, v)
		//每放入一个就进行判断当前最后入栈的是否和判断slice最后一个相等，相等的话就开始判断
		for len(temp) != 0 && temp[len(temp) - 1] == popped[i] {
			//每遇到一个相同的就弹出
			temp = temp[:len(temp) - 1]
			//取出一个就+1
			i++
		}
	}
	//如果最后弹出个数和判断序列相同，就说明是对的
	if i == len(pushed) {
		return true
	}
	return false
}