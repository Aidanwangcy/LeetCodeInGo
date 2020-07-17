package main

import (
	"fmt"
)

func main()  {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	ans := majorityElement(nums)
	fmt.Println(ans)
}

func majorityElement(nums []int) int {
	start := 0
	end := len(nums) - 1
	length := 0
	if len(nums) % 2 == 0 {
		length = len(nums) / 2
	}
	if len(nums) % 2 == 1 {
		length = (len(nums) - 1) / 2
	}
	index := qsort(nums, start, end)
	for index != length {
		if index > length {
			end = index - 1
			index = qsort(nums, start, end)
		} else {
			start = index + 1
			index = qsort(nums, start, end)
		}
	}
	return nums[index]
}

func qsort(nums []int, left, right int) int {
	temp := nums[left]
	for left < right {
		for left < right && nums[right] > temp {
			right --
		}
		nums[left] = nums[right]
		for left < right && nums[left] < temp {
			left ++
		}
		nums[right] = nums[left]
	}
	nums[left] = temp
	return left
}