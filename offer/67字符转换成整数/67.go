package strtoint

import (
	"math"
	"strings"
)

func StrToInt(str string) int {
	num, i := 0, 0
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	sign := 1 //负号标志
	signM := map[byte]int{'+': 1, '-': -1}
	if str[i] == '+' || str[i] == '-' {
		sign = signM[str[i]]
		i++
	} else if !('0' <= str[i] && str[i] <= '9') {
		return 0
	}

	for ; i < len(str) && '0' <= str[i] && str[i] <= '9'; i++ {
		num = num*10 + int(str[i]-'0')
		// 累加可能越界
		if sign == 1 && num > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && -num < math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * num
}
