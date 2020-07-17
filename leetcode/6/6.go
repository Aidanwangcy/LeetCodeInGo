package main

import (
	"fmt"
)

func main() {
	var s string
	var numRows int
	fmt.Printf("输入字符串:")
	fmt.Scan(&s)
	fmt.Printf("输入行数:")
	fmt.Scan(&numRows)
	ans := convert(s, numRows)
	fmt.Printf("OutPut:%v\n", ans)
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	var res string
	//新建n个字符串切片
	list := make([]string, numRows)
	//移动方向（用来改变切片增减）
	next := 1
	//所在切片标记
	step := 0
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		list[step] += str
		//当step所标记的切片到达n-1或者step为0且next不是向下时进行方向转换
		if step == numRows-1 || (step == 0 && next != 1) {
			next *= -1
		}
		//按照next=1（方向向下）或按照next=-1（方向向上）变动切片
		step += next
	}

	//把切片转换为字符串
	for i := 0; i < len(list); i++ {
		res += list[i]
	}
	return res
}
