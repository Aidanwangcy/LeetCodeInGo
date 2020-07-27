package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Printf("输入n：")
	fmt.Scan(&n)
	ans := countDigitOne(n)
	ans2 := myCountDigitOne(n)
	fmt.Println(ans)
	fmt.Println(ans2)
}

func countDigitOne(n int) int {
	return 0
}

//超时了
func myCountDigitOne(n int) int {
	count := 0
	var ss string
	for i := 1; i <= n; i++ {
		ss = strconv.Itoa(i)
		for _, v := range ss {
			if v == '1' {
				count++
			}
		}
	}
	return count
}
