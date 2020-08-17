package backpackfull

//给定 n 种物品和一个容量为 C 的背包，物品 i 的重量是 wi，其价值为 vi，
//每个物品都有无限多件，现在往背包里面装东西，怎么装能使背包的内物品价值最大？

func backpack_full(nums [][]int, total int) int  {
	dp := make([]int, total+1)

	for i := 0; i < len(nums); i++ {
		for j := nums[i][0]; j <= total; j++ {
			dp[j] = max(dp[j], dp[j - nums[i][0]] + nums[i][1])
		}
	}
	return dp[total]
}

func max(i,j int) int {
	if i>=j {
		return i
	}
	return j
}