package main

import "fmt"

func main() {
	nums := []int{2,3,2}
	ans := rob(nums)
	fmt.Println(ans)
}


func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    } else if n == 2 {
        return Max(nums[0], nums[1])
	}
	//围成环时，要么第一个房间被偷要么最后一个房间被偷，二者不能同时进行
	//要不从0抢到-2，要不从1抢到-1
    return Max(startRob(nums, 0, n-2), startRob(nums, 1, n-1))
}
// 从start到end开始
func startRob(nums []int, start, end int) int {
    pre, max := 0, 0
    for i := start; i <= end; i++ {
        pre, max = max, Max(pre+nums[i], max)
    }
    return max
}
func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}