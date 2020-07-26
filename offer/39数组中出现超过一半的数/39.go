package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	ans := majorityElement(nums)
	ans2 := majorityElement2(nums)
	fmt.Println(ans)
	fmt.Println(ans2)
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//majorityElement2 摩尔投票法
//发生正负抵消时，剩余数组的众数一定不变，因为抵消的数中有一半是x为众数
func majorityElement2(nums []int) int {
	x, votes := nums[0], 1 // x 为假设的众数，votes为票数
	for i := 1; i < len(nums); i++ {
		if votes == 0 { // 票数归零，重新设当前元素为众数，准备开始新一轮的计票
			x = nums[i] 
		}

		if nums[i] == x { // 若当前数字等于众数x，则票数+1，否则-1
			votes++
		} else {
			votes--
		}
	}

	return x
}
