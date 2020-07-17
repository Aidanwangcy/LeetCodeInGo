package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("输入绳子长度：")
	fmt.Scan(&n)
	ans := cuttingRope(n)
	fmt.Printf("最大乘积：%v \n", ans)
}

func cuttingRope(n int) int {
	return 0
}
