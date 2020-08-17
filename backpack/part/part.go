package backpackpart

//给定 n 种物品和一个容量为 C 的背包，物品 i 的重量是 wi，其价值为 vi，每个物品都有ki件
//现在往背包里面装东西，怎么装能使背包的内物品价值最大？
//一个比较简单但是效率不高的方法： 把物品i的ki件物品转化为k件不同的物品，直接就转化为01背包问题了。不过效率较低。
//另一种思路：再申请一个数组来保存物品i的使用个数，从而控制物品i的数量,用完全背包的思路解决问题

func backpack_part(nums [][]int, total int) int {
	//nums[i][0] 物品重量
	//nums[i][1] 物品价值
	//nums[i][2] 物品最大数量

	//背包容量为i时的最大价值
	dp := make([]int, total+1)

	//背包容量为i时放了的物品数量
	count := make([]int, total+1)

	for i := 0; i < len(nums); i++ {
		//重置count，每次循环表示一个物品存放个数
		count = make([]int, total+1)
		for j := nums[i][0]; j <= total; j++ {
			//count[j-nums[i][0]]表示已经放了物品i的个数，如果大于物品i的总个数，跳过
			if count[j-nums[i][0]] < nums[i][2] {
				dp[j] = max(dp[j], dp[j-nums[i][0]]+nums[i][1])
				//如果放入物品i的价值更大，
				//则count[j]需要在count[j-nums[i][0]](之前放入物品的数量)
				//加一
				if dp[j] == (dp[j-nums[i][0]] + nums[i][1]) {
					count[j] = count[j-nums[i][0]] + 1
				}
			}

			if dp[j] ==0 || dp[j] <dp[j-1] {
				dp[j] = dp[j-1]
				count[j] = count[j-1]
			}
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
