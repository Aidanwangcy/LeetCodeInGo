package main

import "fmt"

func main()  {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	ans := maxSlidingWindow(nums, k)
	for _, v := range ans {
		fmt.Print(v, " ")
	}
}

func maxSlidingWindow(nums []int, k int) []int {
    left := 0
    right := left+k-1
    ans := []int{}
    for right <= len(nums)-1 {
		maxInt := nums[left]
		for i:=left; i<=right; i++{
			if nums[i] > maxInt {
				maxInt = nums[i]
			}
		}
        ans = append(ans, maxInt)
        left++
        right++
	}
	return ans
}