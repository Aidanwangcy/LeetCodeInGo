package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	//go tool trace trace.out 运行后使用这条命令查看go gc过程
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	//////////////////////////////
	var nums = []int{0, 2, 3, 4, 5, 6, 7, 9}
	ans := missingNumber(nums)
	fmt.Println(ans)
}

func missingNumber(nums []int) int {
	return binarySearch(nums, 0, len(nums)-1)
}

func binarySearch(nums []int, left, right int) int {
	for left <= right {
		var mid int = (left + right) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			//当取到左半部分的时候0～right
			right = mid - 1
		}
	}
	return left
}
