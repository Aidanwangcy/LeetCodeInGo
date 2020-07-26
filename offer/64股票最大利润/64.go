package main

import (
	"fmt"
)

func main()  {
	ss := []int{7,1,5,3,6,4}
	ans := maxProfit(ss)
	fmt.Printf("%v\n", ans)
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1],prices[i] - min)
		if prices[i] < min {
			min = prices[i]
		}
	}
	return dp[len(prices)-1]
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return  b
}
