package main

import "fmt"

func main() {
	nums := []int{2, 8, 4, 5, 6, 9, 3, 1}
	ans := myQsort(nums)
	for k, v := range ans {
		fmt.Printf("index:%v, value:%v\n", k, v)
	}
}

func myQsort(nums []int) []int {
	return qsort(nums, 0, len(nums)-1)
}

func qsort(nums []int, left, right int) []int {
	if left<right {
		p := partion(nums, left, right)
		qsort(nums, left, p-1)
		qsort(nums, p+1, right)
	}
	return nums
}

func partion (nums []int, left, right int) int {
	p := left
	index := p+1
	for i := index; i <= right; i++ {
		if nums[i] < nums[p] {
			swap(nums, i, index)
			index++
		}
	}
	swap(nums, p, index-1)
	return index-1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
