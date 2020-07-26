package main

import "fmt"

func main() {
	s := "abcabcbb"
	ans1 := lengthOfLongestSubstring1(s)
	ans2 := lengthOfLongestSubstring2(s)
	ans3 := lengthOfLongestSubstring3(s)
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//lengthOfLongestSubstring1 滑动双指针, 时间复杂度o(n2)
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	win := make(map[byte]int)

	left := 0
	right := 0
	ans := 1

	for right < len(s) {
		c := s[right]
		right++
		win[c]++

		//缩小窗口
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		ans = max(ans, right-left)
	}
	return ans
}

//lengthOfLongestSubstring2 map滑动双指针, 时间复杂度o(n)
func lengthOfLongestSubstring2(s string) int {
	res := 0
	left, right := 0, 0
	resMap := make(map[byte]int, 0)
	for right < len(s) {
		_, ok := resMap[s[right]]
		for ok {
			resMap[s[left]]--
			if resMap[s[left]] == 0 {
				delete(resMap, s[left])
			}
			left++
			_, ok = resMap[s[right]]
		}
		resMap[s[right]]++
		if right-left+1 > res {
			res = right - left + 1
		}
		right++
	}
	return res

	// if len(s) < 2 {
	// 	return len(s)
	// }
	// left, right, ans := 0, 0, 0
	// res := make(map[string]int, 0)
	// for i := 0; i < len(s); i++ {
	// 	v, ok := res[string(s[right])]
	// 	if ok {
	// 		left = v + 1
	// 		delete(res, string(s[right]))
	// 	}
	// 	res[string(s[right])] = i
	// 	ans = max(ans, right - left + 1)
	// 	right++
	// }
	// return ans
}

//lengthOfLongestSubstring3 动态规划（线性遍历）
func lengthOfLongestSubstring3(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	res, tmp, i := 0, 0, 0
	for j := 0; j < len(s); j++ {
		i = j - 1
		for i >= 0 && s[i] != s[j] {
			i-- //向左遍历找到相等的一位
		}
		if tmp < j-i {
			tmp++ //如果在上次重复长度之外，仅加一
		} else {
			tmp = j - i //反之等于这次的长度
		}
		res = max(res, tmp) //选择最大值
	}
	return res
}
