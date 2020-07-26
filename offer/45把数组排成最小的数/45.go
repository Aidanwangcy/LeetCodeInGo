package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{10, 2}
	ans := minNumber(nums)
	fmt.Println(ans)
}

type Slice []int

func (p Slice) Len() int {
	return len(p)
}

func (p Slice) Less(i, j int) bool {
	istr := strconv.Itoa(p[i])
	jstr := strconv.Itoa(p[j])

	s1 := istr + jstr
	s2 := jstr + istr

	if s1 < s2 {
		return true
	}
	return false
}

func (p Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func minNumber(nums []int) string {
	var snums Slice = Slice(nums)

	sort.Sort(snums)

	var ans string

	for _, v := range snums {
		ans += strconv.Itoa(v)
	}
	return ans
}
