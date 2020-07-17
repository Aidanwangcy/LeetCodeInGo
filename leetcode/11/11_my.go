package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	var heights string
	fmt.Print("输入高度(大于两个,用逗号隔开)：")
	fmt.Scan(&heights)
	s := strings.Split(heights, ",")
	var height []int = make([]int, len(s))
	for i, v := range s {
		height[i], _ = strconv.Atoi(v)
	}
	ans := maxArea(height)
	fmt.Printf("水量：%v\n", ans)
}

func maxArea(height []int) int {
	l := len(height)
	var h int
	maxV := 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if height[i] <= height[j] {
				h = height[i]
			}
            if height[j] < height[i] {
				h = height[j]
			}
			t := j-i
			v := h * t
			if v > maxV {
				maxV = v
			}
		}
	}
	return maxV	
}