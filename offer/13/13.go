package main

import "fmt"

func main()  {
	m, n, k := 2,3,1
	ans := movingCount(m,n,k)
	ans2 := movingCount2(m,n,k)
	fmt.Println(ans, " ", ans2)
}

func movingCount(m,n,k int) int {
	array := make([][]int, m)
	for i, _ := range array {
		array[i] = make([]int, n)
	}

	array[0][0] = 1
	a := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i>0 && array[i-1][j] == 1) || (j>0 && array[i][j-1] == 1) {
				if getres(i, j, k) {
					array[i][j] = 1
					a++
				}
			}
		}
	}
	return a
}

func getres(m, n, k int) bool {
	res := 0
	for m > 0 {
		res += m%10
		m /= 10
	}
	for n > 0 {
		res += n%10
		n /= 10
	}
	return res <= k
}

func movingCount2(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, i, j, k int, dp [][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] == 1 || (sumPos(i)+sumPos(j)) > k {
		return 0
	}

	dp[i][j] = 1

	sum := 1
	sum += dfs(m, n, i, j+1, k, dp)
	sum += dfs(m, n, i, j-1, k, dp)
	sum += dfs(m, n, i+1, j, k, dp)
	sum += dfs(m, n, i-1, j, k, dp)
	return sum
}

// 求所有位之和
func sumPos(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}