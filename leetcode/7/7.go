package main

import (
	"fmt"
)


func main()  {

	var i int
	fmt.Printf("Input：")
	fmt.Scan(&i)
	ans := reverse(i)
	fmt.Printf("OutPut:%v\n", ans)
}

// 1. 0x80 00 00 00 == 1 << 31,即-2147483648（最小负数）
// 2, 原码与补码之和为0，也就是a + ~a + 1 = 0. 所以 -a = 1+ ~a, -（0x80000000) = 1 + 0x7fffffff（最大正数） = 0x80000000
// 3. 1 << 31 是 -2147483648,最小负数. -(-2147483648) = 2147483647 + 1 是 最大正数 + 1 = 最小负数 = -2147483648 
//    其实问题等价于 为什么最大的正数 + 1 变成最小的负数。

func reverse(x int) int {
	y := 0
	for x!=0 {
		//交换：取模后加到对应的位置之上
		y = y*10 + x%10
		//判断越界
		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}