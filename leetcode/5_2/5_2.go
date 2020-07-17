package main

import "fmt"

func main() {
	var s,ans string
	fmt.Println("请输入字符串：")
	fmt.Scanln(&s)
	ans = longestPalindrome(s)
	fmt.Println(ans)
}

func expandAroundCenter(s string, left, right int) (int, int){
	for ; left >= 0 && right <len(s) && s[left] == s[right]; left, right = left - 1, right+1 {}
	//循环从left，right位置向两边扩展，知道找到左右相同的两个索引值
	return left+1, right-1
}

func longestPalindrome(s string) string {
	if s == ""{
		return s
		//空字符串直接返回
	}
	start, end := 0,0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		//从每一个字符开始，对应存在字符长度为1及以上的情况
		left2, right2 := expandAroundCenter(s, i, i+1)
		//从两个字符开始，对应存在字符长度为2及以上的情况
		if right1 - left1 > end - start{
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
		//以上两个挑选最长的两个角标
	}
	return s[start:end+1]
}

