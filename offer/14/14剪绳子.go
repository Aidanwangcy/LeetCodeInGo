package main

import (
	"fmt"
	"math"
)

func main()  {
	var n int
	fmt.Print("输入绳子长度：")
	fmt.Scan(&n)
	ans := cuttingRope(n)
	ans2 := cuttingRope2(n)
	fmt.Printf("最大乘积：%v or %v\n", ans, ans2)
}

//动态规划
func cuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1
	//
	for i := 2; i < n+1; i++ {
		j, k := 1, i-1
		res := 0
		for j <= k {
			res = max(res, max(j, dp[j]) * max(k, dp[k]))
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}

//贪心算法
func cuttingRope2(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	a := math.Floor(float64(n) / 3)
	b := n % 3

	if b == 0 {
		return int(math.Pow(3, a))
	}

	if b == 1 {
		return int(math.Pow(3, a-1)*4)
	}

	return int(math.Pow(3, a) * 2)

}