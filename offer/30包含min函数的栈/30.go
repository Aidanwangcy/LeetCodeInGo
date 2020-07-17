package main

import "fmt"

type MinStack struct{
	nums []int
	min int
}

func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

func (this *MinStack) Push(x int) {
	if len(this.nums) == 0 {
		this.min = x
	} else {
		if x < this.min {
			this.min = x
		}
	}
	this.nums = append(this.nums, x)
}

func (this *MinStack) Pop() {
	temp := this.nums[len(this.nums) - 1]
	this.nums = this.nums[0 : len(this.nums) - 1]
	if temp == this.min {
		if len(this.nums) > 0 {
			this.min = this.nums[0]
			for _, v := range this.nums {
				if v < this.min {
					this.min = v
				}
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums) - 1]
}

func (this *MinStack) Min() int {
	return this.min	
}

func main()  {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	a := obj.Min()
	obj.Pop()
	b := obj.Top()
	c := obj.Min()
	fmt.Printf("%v, %v, %v\n",a,b,c)
}