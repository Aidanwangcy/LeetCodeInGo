package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	for _, v := range ans {
		for _, k := range v {
			fmt.Printf("%v ", int(k))
		}
		fmt.Printf("\n")
	}
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	//第一个数从第一位开始
	for first := 0; first < n; first++ {
		//判断当前数不能和前一个数相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//第三个数取最后一个
		third := n - 1
		//取反
		target := -1 * nums[first]
		//第二个数从第一个数的后一个开始
		for second := first + 1; second < n; second++ {
			//判断当前数不能和前一个数相同
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			//两数还没重合，但是和已经大于0，向内移动第三个数，直到两个位置满足或相等
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			//相等就跳出
			if second == third {
				break
			}
			//如果为0就添加
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
