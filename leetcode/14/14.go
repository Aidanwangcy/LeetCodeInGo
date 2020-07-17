package main

import (
	"fmt"
	"strings"
)

func main() {
	var strs string
	fmt.Print("输入字符串（用逗号隔开）：")
	fmt.Scan(&strs)
	str := strings.Split(strs, ",")
	ans := longestCommonPrefix(str)
	fmt.Println(ans)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	//注意i从1开始，也就是从第二个开始比较
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	lenght := min(len(str1), len(str2))
	index := 0
	//如果第一个字母就不相同返回空字符串
	//切片使用[:2]表示时，左闭右开，若为[:0]则为空
	for index < lenght && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
