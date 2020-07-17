package main

import (
	"fmt"
)

func main() {
	var s, p string
	fmt.Print("主串：")
	fmt.Scan(&s)
	fmt.Print("模式串：")
	fmt.Scan(&p)
	ans := isMatch(s, p)
	fmt.Printf("是否匹配：%v\n", ans)
}

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls+1)
	////////////////////////////////
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, lp + 1)
	}
	///////////////////////////////


	dp[0][0] = true
	for j := 1; j <= lp; j++ {
		if p[j-1] == '*' && dp[0][j-2] {
			dp[0][j] = true
		}
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			nows := s[i-1]
			nowp := p[j-1]
			if nows == nowp {
				dp[i][j] = dp[i-1][j-1]
			} else if nowp == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if nowp == '*' {
				if j >= 2 {
					nowpLast := p[j-2]
					if nowpLast == nows || nowpLast == '.' {
						//p[j-2]为普通字符或点
						dp[i][j] = dp[i-1][j] || dp[i][j-1]
					}
					dp[i][j] = dp[i][j-2] || dp[i][j]
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[ls][lp]
}
