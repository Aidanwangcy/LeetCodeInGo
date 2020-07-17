package main

import "fmt"

func main() {
	var s,ans string
	fmt.Println("请输入字符串：")
	fmt.Scanln(&s)
	ans = longestPalindrome(s)
	fmt.Println(ans)
}

func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int , n)
	//第一行填充字符串
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	
	for l := 0; l < n; l++ {
		for i := 0; i + l < n; i++ {
			//i是字符串中最前面一个
			j := i + l
			//j是字符串中最后面一个
			if l == 0 {
				dp[i][j] = 1
				//当字符串长度为1的时候肯定是回文串
			} else if l == 1 {
				//当长度为2时
				if s[i] == s[j] {
					dp[i][j] = 1
					//如果两个字符相等，肯定是回文串
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
					//状态转移方程，如果当前两个相等，且字串是回文串，当前串就也是回文串
					//只有 s[i+1:j-1]s[i+1:j−1] 是回文串，并且 s 的第 i 和 j 个字母相同时，s[i:j]s[i:j] 才会是回文串
					//从长度短的字符串，转移到长度长的
				}
			}
			if dp[i][j] > 0 && l + 1 > len(ans) {
				ans = s[i:i+l+1]
				//dp为1，且长度最大的字串
			}
		}
	}
	return ans
}