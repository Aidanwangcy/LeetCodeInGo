package backpack01

func backpack01(nums [][]int, total int) int {
	//总结：背包问题三步走：
	//（1）初始化dp数组，行为物品个数+1，列为总重量+1
	//（2）初始化边界，只放一个物品，在不同总重量下得到的价值
	//（3）遍历数组，依赖dp[i-1]更新dp[i]

	//eg.物品分别为{5，12}， {4，3}，{7，10}，{2，3}， {6，6}，背包总容量为15
	//nums[i][0]物品重量
	//nums[i][1]物品价值
	//total背包重量
	
	//dp[i][j]表示总重量为j，物品种类为i，可以获得的最大价值
	dp := make([][]int, len(nums))
	//列号代表背包容量
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total+1)
	}

	//初始化第一行，放入第一个物品，从容量为第一个物品体积开始nums[0][0]，到背包总重量，只放第一个物品，最大价值都为第一个物品价值
	for i := nums[0][0]; i < total; i++ {
		dp[0][i] = nums[0][1]
	}

	//遍历每一个物品
	for i := 1; i < len(nums); i++ {
		//计算背包容量为w时的最大值
		for w := nums[i][0]; w <= total; w++ {
			dp[i][w] = max(dp[i-1][w], dp[i-1][w-nums[i][0]]+nums[i][1])
		}
	}
	return dp[len(nums)-1][total]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
