package main

import "fmt"

func main()  {
	ss := []int{-2,1,-3,4,-1,2,1,-5,4}
	ans := maxSubArray(ss)
	fmt.Println(ans)
}

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    //两个变量，最大值和求和
    max, sum:= nums[0], 0
    for i := 0; i < len(nums); i++ {
        //前面的贡献值为负
        if sum + nums[i] < nums[i] {
            sum = nums[i]
        } else {//nums[i]贡献值为正
            sum += nums[i]
        }
        //记录最大值
        if sum > max {
            max = sum
        }
    }
    return max
}