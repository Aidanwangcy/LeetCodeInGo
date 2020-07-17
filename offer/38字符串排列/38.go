package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Print("输入一个字符串：")
	fmt.Scan(&s)
	ans := permutation(s)
	for _, v := range ans {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}

func permutation(s string) []string {
    if s == "" || len(s) == 0 {
        return []string{}
	}
	var ans []string
	helper([]byte(s), 0, &ans)
	return ans
}

func helper(s []byte, start int, ans *[]string) {
	if start == len(s) {
		*ans = append(*ans, string(s))
	}
	m := make(map[byte]int)
	for i := start; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		s[i], s[start] = s[start], s[i]
		helper(s, start + 1, ans)
		s[i], s[start] = s[start], s[i]
		m[s[i]] = 1
	}
}
