package mymergesort

import "testing"

func TestMertgeSort(t *testing.T)  {
	value := []int{1, 4, 9, 6, 7, 8, 5, 2, 3}
	value2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans := MergeSort(value)
	if len(ans) != len(value2) {
		t.Errorf("MergeSort() fail, len() expect %v, got %v", len(value2), len(ans))
	}
	for i := 0; i < len(ans); i++ {
		if ans[i] != value2[i] {
			t.Errorf("MergeSort() fail, expect %v , got %v", value2[i], ans[i])
		}
	}
}