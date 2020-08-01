package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 1, 2, 6}
	ans := isStraight(nums)
	fmt.Println(ans)
}

func isStraight(nums []int) bool {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			count++
			continue
		}
		dif := nums[i+1] - nums[i]
		count = count - dif + 1
		if count < 0 || dif == 0 {
			return false
		}
	}
	return true
}
