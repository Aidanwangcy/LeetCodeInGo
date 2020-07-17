package main

import (
	"fmt"
	"math"
	"strings"
)

func main()  {
	var str string
	fmt.Printf("InPut:")
	fmt.Scan(&str)
	ans := atoi(str)
	fmt.Printf("OutPut:%v\n", ans)
}

func atoi(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i,v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}

		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign*result
}