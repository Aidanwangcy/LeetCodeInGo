package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue"
	ans := reverseWords(s)
	fmt.Println(ans)
}

func reverseWords(s string) string {
	//使用空格切分字符串到切片中
	stringList := strings.Split(s, " ")
	var ans []string
	for i := len(stringList) - 1; i >= 0; i-- {
		//去掉字符串中的空格
		str := strings.TrimSpace(stringList[i])
		if len(str) > 0 {
			ans = append(ans, str)
		}
	}
	//使用空格连接字符串
	return strings.Join(ans, " ")
}
