package myqsort

import "testing"

func TestQsortInt(t *testing.T) {
	value := []int{1, 4, 9, 6, 7, 8, 5, 2, 3}
	value2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans := QsortInt(value)
	for i := 0; i < len(ans); i++ {
		if ans[i] != value2[i] {
			t.Errorf("QsortInt() fail, got %v ,Expect %v",ans[i], value2[i])
		}
	}
}
