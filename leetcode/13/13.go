package main

import "fmt"

func main()  {
	var s string
	fmt.Print("输入一个罗马数字：")
	fmt.Scan(&s)
	ans := romanToInt(s)
	fmt.Printf("outPut:%v\n", ans)
}

func romanToInt(s string) int {
	l := len(s)
	ss := map[int]string{}
	for i, c := range s {
		ss[i] = string(c)
	}

	v := 0

	for i := 0; i < l; i++ {
		if ss[i] == "I" {
			if ss[i+1] == "V" || ss[i+1] == "X" {
				v--
			} else {
				v++
			}
		} else if ss[i] == "X" {
			if ss[i+1] == "L" || ss[i+1] == "C" {
				v -= 10
			} else {
				v += 10
			}
		} else if ss[i] == "C" {
			if ss[i+1] == "D" || ss[i+1] == "M" {
				v -= 100
			} else {
				v += 100
			}
		} else if ss[i] == "V" {
			v += 5
		} else if ss[i] == "L" {
			v += 50
		} else if ss[i] == "D" {
			v += 500
		} else if ss[i] == "M" {
			v += 1000
		}
	}

	return v
}