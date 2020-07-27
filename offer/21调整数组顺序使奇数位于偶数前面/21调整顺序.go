package main

import (
	"fmt"
)

func main()  {
	nums := []int{1,2,3,4,5,6,7,8}
	ans := exchange(nums)
	for _, v := range ans {
		fmt.Print(v)
	}
	fmt.Print("\n")
}

func exchange(nums []int) []int {
	l := len(nums)
	if  l == 0 {
		return nums
	}
	i, j := 0, l-1
	for i != j {
		if nums[i]%2 == 0 {
			if nums[j]%2 == 1 {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return nums
}