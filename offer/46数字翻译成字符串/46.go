package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 12258
	ans := translateNum(num)
	fmt.Println(ans)
}

func translateNum(num int) int {
	ss := strconv.Itoa(num)
	l := len(ss)
	dp := make([]int, l+1)
	dp[0],dp[1] = 1, 1
	for i := 2; i <= l; i++ {
		dp[i] = dp[i-1]
		if temp, _ := strconv.Atoi(ss[i-2:i]); temp<26 && temp>9 {
			dp[i] = dp[i-2] + dp[i-1]
		}
	}
	return dp[l]
}
